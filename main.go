package main

import (
	"database/sql"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/sqltocsv"

	_ "modernc.org/sqlite"
)

func main() {
	var sqlField string

	// accept file name from CLI
	flag.StringVar(&sqlField, "file name", "", `Put file youd like to convert to csvfile`)
	flag.Parse()

	// open database
	db, err := sql.Open("sqlite", "./"+sqlField)
	if err != nil {
		log.Fatal(err)
	}
	// close data base
	defer db.Close()

	// select all data from a table
	result, err := db.Query(`Select * from test`)
	if err != nil {
		log.Fatal(err)
	}
	// reflect response from database to csv
	csvConverter := sqltocsv.New(result)
	// make csv file
	csvConverter.WriteFile("result.csv")
	//  read csv file
	// open csv file
	file, err := os.Open("./result.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read opened file as a csvfile
	reader := csv.NewReader(file)

	var line []string

	for {
		line, err = reader.Read()
		if err != nil {
			break
		}
		fmt.Printf("%v", line)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
