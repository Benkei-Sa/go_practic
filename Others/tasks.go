/*package main

import (
	"fmt"
	"math"
)

func gep(a, b int) int {
	a = a * a
	b = b * b
	c := math.Sqrt((float64(a) + float64(b)))
	return int(c)
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b)
	c = gep(a, b)
	fmt.Print(c)
}
*/

/*package main

import (
	"fmt"
	"strings"
)

func conver(str string) string {
	var result string
	result = strings.Replace(str, "", "*", -1)
	result = strings.Trim(result, "*")
	return result
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Print(conver(str))
}
*/

/*package main

import (
	"fmt"
	"strings"
)

func max_in_str(str string) string {
	temp := strings.Split(str, "")
	max := temp[0]
	for _, v := range temp {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Print(max_in_str(str))

}
*/

/*package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convert_mas(str string) []int {
	temp_str := strings.Split(str, "")
	var int_mas = []int{}
	for _, v := range temp_str {
		mas, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		int_mas = append(int_mas, mas)
	}
	return int_mas
}

func qr_mas_on_index(int_mas []int) {
	for _, v := range int_mas {
		x := v * v
		fmt.Print(x)
	}
}

func main() {
	var str string
	fmt.Scan(&str)
	qr_mas_on_index((convert_mas(str)))
}
*/

func M() int {
	m := p * v
	return m
}

func W() int {
	w := k / (M())
	w = math.Sqrt(w)
	return w
}

func T() int {
	t := 6 / (W())
	return t
}