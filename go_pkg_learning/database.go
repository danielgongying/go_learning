package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	db,err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/shangwei")

	if err != nil {
		log.Fatal(err)
	}

	regmail := "test@test.com"

	var userid uint64
	var nickname string
	var locktime string

	err = db.QueryRow("select UserID, " +
		"RegMail, NickName, LockTime from User where RegMail=? and LockTime>?", regmail, time.Now()).Scan(&userid, &regmail, &nickname, &locktime)

	switch  {
	case err == sql.ErrNoRows:
		log.Fatal(err)
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("UserID: %d \n" +
			"RegMail: %s\nNickName: %s\nLockTime: %s\n",
			userid, regmail, nickname, locktime)
	}
}
