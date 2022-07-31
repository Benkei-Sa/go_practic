package main

import "fmt"

func main() {
	var i float64
	fmt.Scan(&i)
	if i <= 0 {
		fmt.Printf("число %2.2f не подходит", i)
	} else if i > 10000 {
		fmt.Printf("%e", i)
	} else {
		i = i * i
		fmt.Printf("%.4f", i)
	}
}
