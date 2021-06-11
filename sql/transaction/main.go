package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "phpmyadmin:root@/godatabase")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into users(id, name) values(?,?)")

	stmt.Exec(2000, "cleiton rasta")
	stmt.Exec(2001, "cabeça de gelo")
	// _, err = stmt.Exec(1, "otário")

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()
}
