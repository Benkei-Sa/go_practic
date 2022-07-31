package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)
	workArray := [n]uint8{}

	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		workArray[i] = x
	}
	fmt.Print(workArray[3], " ")
}
