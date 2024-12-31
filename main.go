package main

import (
	"bytes"
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

	if !slices.Contains(acceptedOperations, *op) {
		return ErrOperationNotValid
	}

	return nil
}

func createTask(content string) error {
	file, err := openTasks("tasks.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(content + "\n"))
	if err != nil {
		return err
	}

	return nil
}

func readTasks() (string, error) {
	file, err := openTasks("tasks.txt")
	if err != nil {
		return "", nil
	}
	defer file.Close()

	var content bytes.Buffer
	_, err = content.ReadFrom(file)
	if err != nil {
		return "", err
	}
	return content.String(), nil
}

func openTasks(name string) (*os.File, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		file, err := os.Create(name)
		if err != nil {
			return nil, err
		}
		return file, nil
	}

	return os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0)
}
