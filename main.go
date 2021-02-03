package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./test.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Exec(`SELECT * FROM test`)
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()
	fmt.Printf("result: %s", result)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
