package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrOperationNotValid    = errors.New("this operation is not valid. Please try again with a valid operation: add | delete | read | update | complete")
	ErrOperationNotProvided = errors.New("please provide an operation: add | delete | read | update | complete")
	sqliteFile              = "tasks.db"
)

func main() {
	operation := flag.String("operation", "add", "Operation to perform add | delete | read | update | complete")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 && *operation != "read" {
		log.Fatal(ErrOperationNotProvided)
	}

	err := validateOperation(operation)
	if err != nil {
		log.Fatal(err)
	}

	db, err := openDB(sqliteFile)
	if err != nil {
		fmt.Printf("An error occured when connecting to the database %v", err)
		os.Exit(1)
	}

	switch *operation {
	case "add":
		err = createTask(db, args[0])
		if err != nil {
			log.Fatal(err)
		}
	case "read":
		content, err := readTasks(db)
		if err != nil {
			log.Fatal(err)
		}
		log.Print(content)
	default:
		return
	}
}

func validateOperation(op *string) error {
	if len(*op) == 0 {
		return ErrOperationNotProvided
	}

	acceptedOperations := []string{"add", "delete", "read", "update", "complete"}

	if !contains(acceptedOperations, *op) {
		return ErrOperationNotValid
	}

	return nil
}

func openDB(file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER NOT NULL AUTOINCREMENT RIMARY KEY,
		task TEXT NOT NULL,
		completed INTEGER NOT NULL DEFAULT 0 CHECK (completed = 1 OR completed = 0)
	)
	`

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return db, nil
}
