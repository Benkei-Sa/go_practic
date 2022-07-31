package main

import "fmt"

func main() {
	var sum1, sum2, l1, l2, x1, x2 int
	fmt.Scan(&sum1, &sum2)
	n1 := sum1
	for l1 = 0; n1 != 0; l1++ {
		n1 = n1 / 10
	}
	n2 := sum2
	for l2 = 0; n2 != 0; l2++ {
		n2 = n2 / 10
	}
	n1 = sum1
	n2 = sum2
	for l1 != 0 {
		x1 = n1 % 10
		n2 = sum2
		for w := l2; w != 0; w-- {
			x2 = n2 % 10
			if x1 == x2 {
				defer fmt.Print(x1, " ")
			}
			n2 = n2 / 10
		}
		l1--
		n1 = n1 / 10
	}
}
