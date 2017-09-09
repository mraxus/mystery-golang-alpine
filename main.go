package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	addr := os.Getenv("DB")
	fmt.Println("Postgres addr: " + addr)

	_, err := sqlx.Connect("postgres", addr)

	if err != nil {
		fmt.Println("Could not connect...")
	} else {
		fmt.Println("Connecting successful")
	}
}
