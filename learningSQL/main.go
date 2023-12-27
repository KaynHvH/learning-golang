package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func createTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        age INTEGER
    );
    `

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table created successfully!")
}

func addUsers(db *sql.DB, name string, age int) {
	query := "INSERT INTO users (name, age) VALUES (?, ?);"

	_, err := db.Exec(query, name, age)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User added successfully")
}

func getUsers(db *sql.DB) {
	query := "SELECT * FROM users;"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int

		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("ID: %d | NAME: %s | AGE: %d\n", id, name, age)
	}
}

func deleteUser(db *sql.DB, id int) {
	query := "DELETE FROM users WHERE id = ?;"

	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted user with id: %d\n", id)
}

func deleteAllUsers(db *sql.DB) {
	query := "DELETE FROM users;"
	resetIDCounter := "DELETE FROM sqlite_sequence WHERE name = 'users';"

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(resetIDCounter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully deleted all users from database and id counter restarted")
}

func addAgeTransaction(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	query := "UPDATE users SET age = age + 1 WHERE age < 18;"
	_, err = tx.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully transactions were made")
}

func main() {
	db, err := sql.Open("sqlite3", "example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connecting to the database successfully")

	createTable(db)
	deleteAllUsers(db)
	addUsers(db, "Karol", 19)
	addUsers(db, "Adam", 7)
	addUsers(db, "Paulina", 25)
	addUsers(db, "Amelia", 12)
	getUsers(db)
	deleteUser(db, 3)
	addAgeTransaction(db)
	getUsers(db)

}
