package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type GlobalIds []struct {
	Id int `json:"global_id"`
}

func main() {
	var globalIds GlobalIds
	var sum int
	sum = 0
	file := "data-20190514T0100.json"
	byteVal, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(byteVal, &globalIds)
	if err != nil {
		panic(err)
	}
	fmt.Print(json.Unmarshal(byteVal, &globalIds))
	/*for i := 0; i < len(globalIds); i++ {
		sum = sum + (globalIds[i].Id)
	}
	fmt.Println(sum)*/
}
