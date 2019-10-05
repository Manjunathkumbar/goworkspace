package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	data := [][]string{
		{"FirstName", "LastName", "Address"},
		{"Manjunath", "K", "BGM"},
		{"Ranganath", "MK", "MYS"},
		{"Chandru", "B", "ETC22222"},
	}

	csvfile, err := os.Create("data.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvfile)

	for _, row := range data {
		_ = csvwriter.Write(row)
	}

	csvwriter.Flush()
	csvfile.Close()
}
