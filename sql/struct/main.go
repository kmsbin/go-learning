package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "user:root@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	exec(db, "CREATE DATABASE IF NOT EXISTS godatabase")
	exec(db, "USE godatabase")
	exec(db, "DROP TABLE IF EXISTS users")
	exec(db, `CREATE TABLE users (
		id integer auto_increment,
		name varchar(160),
		PRIMARY KEY (id)
	)`)
}
