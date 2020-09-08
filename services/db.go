package services

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//ConnectToDb connetcing the program to the database
func ConnectToDb() {
	db, err := sql.Open("mysql",
		"root:omri3ma1@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	} else {
		println("Connected")
	}
	defer db.Close()
}
