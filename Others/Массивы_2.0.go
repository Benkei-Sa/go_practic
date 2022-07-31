package main

import "fmt"

func main() {
	var workArray [10]uint8
	var a, b, t, x uint8
	for i := 0; i < len(workArray); i++ {
		fmt.Scan(&x)
		workArray[i] = x
	}
	for i := 0; i < 3; i++ {
		fmt.Scan(&a, &b)
		t = workArray[a]
		workArray[a] = workArray[b]
		workArray[b] = t
	}
	for i := 0; i < len(workArray); i++ {
		fmt.Print(workArray[i], " ")
	}
}
