package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("task.data.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()
	newReader := csv.NewReader(file)
	newReader.Comma = ';'
	data, err := newReader.ReadAll()
	for i, v := range data[0] {
		if v == "0" {
			fmt.Println(i + 1)
			break
		}
	}
}

///	for i, v := range mass {
//		if v == "0" {
//			fmt.Print(i + 1)
//		}
//	}
//}
