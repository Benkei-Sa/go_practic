package main

import "fmt"

func main() {
	var sum1, sum2, l1, l2, x1, x2 int
	// получение переменных
	fmt.Scan(&sum1, &sum2)
	// вычесление длинны данных
	n1 := sum1
	for l1 = 0; n1 != 0; l1++ {
		n1 = n1 / 10
	}
	n2 := sum2
	for l2 = 0; n2 != 0; l2++ {
		n2 = n2 / 10
	}
	// цикл проверки
	n1 = sum1
	n2 = sum2
	q := l1 - 1
	w := l2
	for l1 != 0 {
		x1 = n1 % 10
		if l1 == 1 {
			n1 = sum1
			n2 = sum2
			for q != 0 {
				x1 = n1 / 10
				n1 = n1 / 10
				q--
				if x1 < 10 {
					for w != 0 {
						x2 = n2 % 10
						if n2 < 10 && x1 == x2 {
							defer fmt.Print(x2, " ")
						} else if x1 == x2 {
							defer fmt.Print(x2, " ")
						}
						l2--
						n2 = n2 / 10
					}
				}

			}
		} else if n1 < 10 {
			x1 = n1
			for l2 != 0 {
				x2 = n2 % 10
				if n2 < 10 && x1 == n2 {
					defer fmt.Print(n2, " ")
				} else if x1 == x2 {
					defer fmt.Print(x2, " ")
				}
				l2--
				n2 = n2 / 10
			}
		} else {
			for l2 != 0 {
				x2 = n2 % 10
				if n2 < 10 && x1 == n2 {
					defer fmt.Print(n2, " ")
				} else if x1 == x2 {
					defer fmt.Print(x2, " ")
				}
				l2--
				n2 = n2 / 10
			}
		}
		l1--
		n1 = n1 / 10
	}
}
