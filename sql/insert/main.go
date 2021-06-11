package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "phpmyadmin:root@/godatabase")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("insert into users(name) values(?)")

	stmt.Exec("kaulindo")
	stmt.Exec("Maria")
	res, _ := stmt.Exec("Alaska")

	id, _ := res.LastInsertId()
	fmt.Println(id)

	lines, _ := res.RowsAffected()
	fmt.Println(lines)

}
