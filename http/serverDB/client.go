package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		user2Id(w, r, id)
	case r.Method == "GET":
		allUsers(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Sorry... :(")
	}
}

func user2Id(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "phpmyadmin:root@/godatabase")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var u user
	db.QueryRow("select id, name from users where id = ?", id).Scan(&u.ID, &u.Name)

	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application-json")
	fmt.Fprintf(w, string(json))
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "phpmyadmin:root@/godatabase")
	if err != nil {
		log.Fatal(err)
	}
	rows, _ := db.Query("select id, name from users")
	defer db.Close()

	var users []user
	for rows.Next() {
		var user user
		rows.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}

	json, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
