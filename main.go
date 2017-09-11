package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	dbname   = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)

	fmt.Println("Postgres connecting: " + psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Could not connect...")
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Could not ping...")
	} else {
		fmt.Println("Connecting successful")
	}
}