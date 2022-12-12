package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	filename := "export.csv"
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)
	m := make(map[string]int)
	lineNumber := 0
	totalDuplicates := 0
	for {
		var lineContentsSb strings.Builder
		lineContents := ""
		record, err := r.Read()

		if err == io.EOF {
			break
		}
		lineNumber++

		if err != nil {
			log.Fatal(err)
		}

		for value := range record {
			lineContentsSb.WriteString(record[value])
		}
		lineContents = lineContentsSb.String()
		duplicateLineNumber, prs := m[lineContents]
		if prs {
			fmt.Println("Duplicated lines:", duplicateLineNumber, "=", lineNumber)
			totalDuplicates++
		} else {
			m[lineContents] = lineNumber
		}
	}
	fmt.Println(totalDuplicates, "duplicates found")
	fmt.Println("Total of lines:", lineNumber)
}
