package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}
	// fmt.Printf("Name: %s\tSize: %d byte\tPath: %s\n", info.Name(), info.Size(), path)
	return search(path)
}

func search(path string) error {
	file, err := os.Open("task.data.txtth")
	if err != nil {
		return err
	}
	defer file.Close()

	newReader := csv.NewReader(file)
	data, err := newReader.ReadAll()

	if len(data) == 10 {
		fmt.Println(data[4][2])
	}
	return nil
}

func main() {
	const root = "task"

	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}
}
