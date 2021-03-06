package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func insert(db *sql.DB) {
	/*执行SQL语句*/
	stmt, err := db.Prepare("INSERT INTO user(username, password) VALUES(?, ?)")
	defer stmt.Close()

	if err != nil {
		log.Println(err)
		return
	}
	stmt.Exec("daniel", "123123")
	stmt.Exec("john", "123123")

}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	insert(db)

	rows, err := db.Query("select id, username from user where id = ?", 1)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()
	var id int
	var name string
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

