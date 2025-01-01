package main

import (
	"errors"
	"flag"
	"log"
)

var (
	ErrOperationNotValid    = errors.New("this operation is not valid. Please try again with a valid operation: add | delete | read | update | complete")
	ErrOperationNotProvided = errors.New("please provide an operation: add | delete | read | update | complete")
)

func main() {
	operation := flag.String("operation", "add", "Operation to perform add | delete | read | update | complete")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		log.Fatal(ErrOperationNotProvided)
	}

	err := validateOperation(operation)
	if err != nil {
		log.Fatal(err)
	}

	switch *operation {
	case "add":
		err = createTask(args[0])
		if err != nil {
			log.Fatal(err)
		}
	case "read":
		content, err := readTasks()
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

	if contains(acceptedOperations, *op) {
		return ErrOperationNotValid
	}

	return nil
}
