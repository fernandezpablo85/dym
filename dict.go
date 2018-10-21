package main

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
	"strconv"

	"github.com/gobuffalo/packr"
)

var (
	// Dict is the string - count dictionary
	Dict = map[string]int{}

	// DictTotal is the total count of occurences in the dictionary
	DictTotal = 0
)

func init() {
	box := packr.NewBox("./assets")
	bs, err := box.MustBytes("dict.csv")
	if err != nil {
		log.Println("warning: main dict file (dict.csv) not found, falling back to dict.csv.example")
		bs, err = box.MustBytes("dict.csv.example")
		if err != nil {
			log.Fatalf("no .csv file found: %v\n", err)
		}
	}

	r := bytes.NewReader(bs)
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading .csv: %v", err)
		}
		name := record[0]
		count, err := strconv.Atoi(record[1])
		Dict[name] = count
		DictTotal += count
	}
}
