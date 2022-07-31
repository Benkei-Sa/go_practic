var a [5]int

a[0] = 8
a[1] = 42

fmt.Println(a[1]) // выведет 42

Срез массива 
var s []int = a[1:3]

//

package main

import "fmt"

func main() {
  a := [5]int{0, 2, 4, 6, 8}
  var s []int = a[1:3]

  s[0] = 8
  fmt.Println(a) // выведет [0 8 4 6 8]
}

//

a := make([]int, 5)
a = append(a, 8)
fmt.Println(a) // выведет [0 0 0 0 0 8]

// перебор массива

package main

import "fmt"

func main() {
  a := make([]int, 5)
  a[1] = 2
  a[2] = 3

  for i, v := range a {
    fmt.Println(i, v)
  }
}

// вывод фактических символов

x := "hello"
for _, c := range x {
  fmt.Printf("%c \n", c)
}

// создание карты 
m := make(map[string]int) 

// примет карты 
m["James"] = 42
m["Amy"] = 24

fmt.Println(m["James"]) // выведет 42

// удалению карты
m := map[string]int{
	"James": 42, 
	"Amy": 24}
	 
 delete(m, "James")

//

func sum(nums ...int) {
	total := 0
	for _, v := range nums {
	  total += v
	}
	fmt.Println(total)
  }

  //
s := []int{42, 33, 22, 11}
sum(s...)
//Без трех точек в конце аргумента вместо нескольких значений
//будет получен сам срез, что приведет к ошибке.

