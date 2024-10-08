package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Post struct {
	id      int
	title   string
	content string
	author  string
}

var conn *pgx.Conn

func main() {
	var err error

	dbUrl := "postgres://postgres:docker@localhost:5432/goweb"
	conn, err = pgx.Connect(context.Background(), dbUrl)

	if err != nil {
		panic(err)
	}

	defer conn.Close(context.Background())

	// createTable()

	// insertPost()

	// insertPostWithReturn()

	selectAllPosts()
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

func insertPost() {
	title := "Post 2"
	content := "Conteúdo do post 2"
	author := "Lucas Silva Leoncio"

	query := fmt.Sprintf(`
		INSERT INTO posts (title, content, author)
		values ('%s', '%s', '%s')
	`, title, content, author) // Metodo errado

	/*
		Utilizando a forma apresentado na query logo acima, abre espaço para injeção de comando maliciosos para o sistema
	*/

	fmt.Println(query)

	new_query := `
		INSERT INTO posts (title, content, author)
		values ($1, $2, $3)
	`

	conn.Exec(context.Background(), new_query, title, content, author)
}

func insertPostWithReturn() {
	title := "Post 5"
	content := "Conteúdo do post 5"
	author := "Lucas Silva Leoncio 5"

	query := `
		INSERT INTO posts (title, content, author)
		values ($1, $2, $3)
		RETURNING id
	`

	row := conn.QueryRow(context.Background(), query, title, content, author)

	var id int

	if err := row.Scan(&id); err != nil {
		panic(err)
	}

	fmt.Println("Post Criado. Id = ", id)
}

func selectPostById() {
	query := `
		select title, content, author from posts where id = $1
	`

	var title, content, author string

	row := conn.QueryRow(context.Background(), query, 3)

	if err := row.Scan(&title, &content, &author); err != nil {
		panic(err)
	}

	fmt.Printf("Post: title:%s, content:%s, author:%s", title, content, author)
}

func selectAllPosts() {
	query := `
		select * from posts
	`

	rows, err := conn.Query(context.Background(), query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Err() != nil {
		panic(rows.Err())
	}

	var posts []Post

	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.id, &post.title, &post.content, &post.author); err != nil {
			panic(err)
		}

		posts = append(posts, post)
	}

	for _, post_ := range posts {
		fmt.Println(post_)
	}
}
