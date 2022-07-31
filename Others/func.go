/* создание функции
func welcome() {
	fmt.Println("Привет")
}

//Вызов происходит только по названию функции
// Чтобы вернуть полученое значение используем return
	func sum(x, y int) int {
  return x + y
}

func calc(a int) (int, int) {
  return a + 2, a + 1 
} 

func main() {
  x, y := calc(5)
  fmt.Println(x + y)
}


// Оператор defer гарантирует, что функция будет вызвана
только после того, как вернется окружающая функция
Повледний пришел, первый вышел, складывается в стек стопкой

func double(a int) {
  fmt.Println(a*2) 
} 

func main() {
  for i := 4; i > 0;i-- {
    defer double(i)
  }
}


*/



package main

import "fmt"

func max(a, b int) (int, int) {
	if a > b {
		fmt.Println(a)
	else {
		fmt.Println(b)
	}
	}
}

func main() {

}



var a, b int
a = 0
b = 0
if (n >= 0 || n <= 0) {
	a = n*2
	b = n*n
}

func double_m (a, b int) (int) {
    var l, x, sum int 
	sum = 0
	if a < b {
		l = b - a + 1
		for z := l; z > 0; z-- {
			x = a * a
			sum = sum + x
			a++
		}
	}
	if b < a {
		l = a - b + 1
		for z := l; z > 0; z-- {
			x = b * b
			sum = sum + x
			b++
		}
	}
	if a = b && a = 0 { // обработка нулевых входных данных
		sum = 0
	} else {	// обработка равных входных данных
		sum = a * a
	}
	return sum
}

func double_m (a, b int) (int) {
    var l, x, sum int
    sum = 0
    l = b - a + 1
    for z:=l; z > 0; z-- {
        x = a * a;
        sum = sum + x;
        a++;
    }
    return sum
}