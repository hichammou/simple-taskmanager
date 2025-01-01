package main

import "os"

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
	content, err := os.ReadFile("tasks.txt")
	if err != nil {
		return "", nil
	}
	return string(content), nil
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
