package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

type (
	Student struct {
		Rating []int
	}

	Group struct {
		Students []Student
	}

	Rating struct {
		Average float32
	}
)

func main() {
	var num_stud int
	var sum, i int
	var result Group
	file, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(file, &result)
	for idx, _ := range result.Students {
		num_stud = idx
	}
	for _, student := range result.Students {
		for i = 0; i < len(student.Rating); i++ {

		}
		sum = sum + i
	}
	num_stud = num_stud + 1
	aver := Rating{(float32(sum) / float32(num_stud))}
	data2, err := json.MarshalIndent(aver, "", "    ")
	io.WriteString(os.Stdout, string(data2))
}
