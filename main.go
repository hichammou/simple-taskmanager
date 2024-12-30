package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"slices"
)

var (
	ErrOperationNotValid    = errors.New("this operation is not valid. Please try again with a valid operation: add | delete | read | update | complete")
	ErrOperationNotProvided = errors.New("please provide an operation: add | delete | read | update | complete")
)

func main() {
	operation := flag.String("operation", "add", "Operation to perform add | delete | read | update | complete")
	flag.Parse()

	args := flag.Args()

	err := validateOperation(operation)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(args)

	switch *operation {
	case "add":
		err = createTask(args[0])
		if err != nil {
			log.Fatal(err)
		}
	default:
		return
	}
}

func validateOperation(op *string) error {
	if len(*op) == 0 {
		return ErrOperationNotProvided
	}

	acceptedOperations := []string{"add", "delete", "read", "update", "complete"}

	log.Print(*op)
	if !slices.Contains(acceptedOperations, *op) {
		return ErrOperationNotValid
	}

	return nil
}

func createTask(content string) error {
	if _, err := os.Stat("tasks.txt"); os.IsNotExist(err) {
		_, err := os.Create("tasks.txt")
		if err != nil {
			return err
		}
	}

	file, err := os.OpenFile("tasks.txt", os.O_APPEND, os.ModeAppend)
	_, err = file.Write([]byte(content + "\n"))
	if err != nil {
		return err
	}

	return nil
}
