package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "gee.db")
	if err != nil {
		log.Fatal("Can't open sqlite:", err)
	}
	defer func() { _ = db.Close() }()

	db.Exec("DROP TABLE IF EXIST User;")
	db.Exec("CREATE TABLE User(Name text);")
	if res, err := db.Exec("Insert into User(`Name`) values (?),(?)", "Tom", "Sam"); err == nil {
		affected, _ := res.RowsAffected()
		log.Println("affected: ", affected)
	} else {
		log.Fatal("Exec failed")
	}

	row := db.QueryRow("SELECT Name FROM User LIMIT 1;")
	var name string
	if err := row.Scan(&name); err != nil {
		log.Fatal("query failed")
	}
	log.Println("name is:", name)

	log.Println("=============")
	rows, _ := db.Query("SELECT Name FROM User;")
	for rows.Next() {
		if err := rows.Scan(&name); err != nil {
			log.Fatal("query failed: ", err)
		}
		log.Println("name is:", name)
	}
}
