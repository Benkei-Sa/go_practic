// w = 3 | l = 0 | d = 1

package main

import "fmt"

func main() {
	var w, l, d, sum int
	var input string
	fmt.Scanln(&input)
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
	results = append(results, input)
	for _, v := range results {
		if v == "w" {
			w += 3
		}
		if v == "l" {
			l += 0
		}
		if v == "d" {
			d += 1
		}
		sum = w + l + d
	}
	fmt.Println(sum)
}
