package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func MySQL() {
	db, _ = sql.Open("mysql", "myuser:myuserpassword@(127.0.0.1:3306)/mydatabase?parseTime=true")


	// Initialize the first connection to the database, to see if everything works correctly.
	// Make sure to check the error.
	err := db.Ping()

	if err != nil {
		fmt.Printf("err: %s\n", err)
		return
	}

	// createTable()
	// insert()
	// selectSingle()
	selectAll()
	delete()
	selectAll()
}

func createTable() {
	query := `
    CREATE TABLE users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`

	// Executes the SQL query in our database. Check err to ensure there was no error.
	_, err := db.Exec(query)
	if err != nil {
		
	}
}

func insert() {
	username := "Juju"
	password := "gostoso"
	createAt := time.Now()

	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createAt)	

	if err != nil {
		fmt.Printf("err: %s\n", err)
	}

	userID, err := result.LastInsertId()

	if err != nil {
		fmt.Printf("err: %s\n", err)
	}

	fmt.Printf("id: %d\n", userID)
}

func selectSingle() {
	var (
    id        int
    username  string
    password  string
    createdAt time.Time
	)

	// Query the database and scan the values into out variables. Don't forget to check for errors.
	query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
	err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt)

	if err != nil {
		fmt.Printf("err: %s\n", err)
		return
	}

	fmt.Printf("id: %d, user: %s, pass: %s, created_at: %s", id, username, password, createdAt)
}

func selectAll() {
	type user struct {
		id int
		username string
		password string
		createdAt time.Time
	}

	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
	if err != nil {
		fmt.Printf("Err: %s", err)
		return
	}

	defer rows.Close()
	
	var users []user
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		if err != nil {
			fmt.Printf("Err: %s", err)
			return
		}
		users = append(users, u)
	}

	err = rows.Err()
	if err != nil {
		fmt.Printf("Err: %s", err)
		return
	}

	for _, user := range users {
			fmt.Printf("id: %d, user: %s, pass: %s, created_at: %s\n", user.id, user.username, user.password, user.createdAt)
	}
}

func delete() {
	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
	if err != nil {
		fmt.Printf("Err: %s", err)
	}
}