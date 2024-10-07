package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func main() {
	var err error

	dbUrl := "postgres://postgres:docker@localhost:5432/goweb"
	conn, err = pgx.Connect(context.Background(), dbUrl)

	if err != nil {
		panic(err)
	}

	defer conn.Close(context.Background())

	createTable()
}

func createTable() {
	query := `
		CREATE TABLE IF NOT EXISTS posts(
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			content TEXT,
			author TEXT NOT NULL
		);
	`

	_, err := conn.Exec(context.Background(), query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Table criada")
}
