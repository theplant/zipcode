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

		data, err := json.Marshal(addr)
		if err != nil {
			panic(fmt.Sprintf("rowNum: %v marshal json with err %v", rowNum, err))
		}

		fileName := filepath.Join(*targetDir, record[0]+".json")
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
