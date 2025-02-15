package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func checkError(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func main() {
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3305)/%v", DBUser, DBPassword, DBName)
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

}
