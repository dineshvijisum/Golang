package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("go mysql driver select query resuts")
	db, err := sql.Open("mysql", "root:root12345@tcp(localhost:3306)/dinesh")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("connected")
	defer db.close()
}