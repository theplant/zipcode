package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"golang.org/x/text/encoding/japanese"
)

var sourceFile = flag.String(
	"source-file",
	"tmp/KEN_ALL_ROME.CSV",
	"csv source file name",
)

var targetDir = flag.String(
	"target-dir",
	"jp",
	"store generated json files",
)

var verbose = flag.Bool(
	"verbose",
	false,
	"log progress",
)

// Address is a japan address in Kanji format
// http://www.post.japanpost.jp/zipcode/dl/readme_ro.html
type Address struct {
	Prefecture string `json:"prefecture"`
	City       string `json:"city"`
	Town       string `json:"town"`
}

var manualFixTowns = []string{
	"野々宿",
	"下左草",
	"小繋沢",
	"湯之沢",
	"槻沢",
	"湯田",
	"種市",
}

func main() {
	flag.Parse()

	f, err := os.Open(*sourceFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	csfr := csv.NewReader(f)
	decoder := japanese.ShiftJIS.NewDecoder()
	var rowNum int
	var addr = &Address{}
	for {
		record, err := csfr.Read()
		rowNum++

		if err != nil {
			if err == io.EOF {
				err = nil
				break
			}

			panic(fmt.Sprintf("rowNum: %v read with err %v", rowNum, err))
		}

		addr.Prefecture, err = decoder.String(record[1])
		if err != nil {
			panic(fmt.Sprintf("rowNum: %v, colNum: %v decode with err %v", rowNum, 1, err))
		}
		addr.City, err = decoder.String(record[2])
		if err != nil {
			panic(fmt.Sprintf("rowNum: %v, colNum: %v decode with err %v", rowNum, 2, err))
		}
		addr.Town, err = decoder.String(record[3])
		if err != nil {
			panic(fmt.Sprintf("rowNum: %v, colNum: %v decode with err %v", rowNum, 3, err))
		}

		var removeSpace = func(s string) string {
			return strings.Map(func(o rune) rune {
				if unicode.IsSpace(o) {
					return -1
				}

				return o
			}, s)
		}

		// https://github.com/theplant/zipcode/issues/2
		addr.City = removeSpace(addr.City)

		// https://github.com/theplant/zipcode/issues/1
		if addr.Town == "以下に掲載がない場合" {
			addr.Town = ""
		}

		// remove （１～４丁目）
		addr.Town = strings.Split(addr.Town, "（")[0]

		// fix things like 種市第１地割～第３地割, 下左草７７地割～下左草８０地割, ２２５７－２８、２３１６～２３１８, －４、５４０７－５、５４４５～５４
		if strings.Index(addr.Town, "～") > 0 || strings.Index(addr.Town, "－") > 0 {
			fixed := false
			for _, prefix := range manualFixTowns {
				if strings.Index(addr.Town, prefix) == 0 {
					addr.Town = prefix
					fixed = true
					break
				}
			}
			if !fixed {
				addr.Town = ""
			}
		}

		// https://github.com/theplant/zipcode/issues/2
		addr.Town = removeSpace(addr.Town)

		fileName := filepath.Join(*targetDir, record[0]+".json")

		oldaddr := &Address{}
		data, err := ioutil.ReadFile(fileName)
		if err == nil {
			err = json.Unmarshal(data, oldaddr)
		}
		if oldaddr.Prefecture == addr.Prefecture &&
			oldaddr.City == addr.City && oldaddr.Town == addr.Town {
			if *verbose {
				fmt.Printf("%s not changed\n", fileName)
			}
			continue
		}

		data, err = json.Marshal(addr)
		if err != nil {
			panic(fmt.Sprintf("rowNum: %v marshal json with err %v", rowNum, err))
		}

		err = ioutil.WriteFile(
			fileName,
			data,
			0644,
		)
		if err != nil {
			panic(fmt.Sprintf("rowNum: %v write json with err %v", rowNum, err))
		}

		if *verbose {
			fmt.Printf("%s: %s\n", fileName, string(data))
		}
	}
}
