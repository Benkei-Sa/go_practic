/*package main

import "fmt"

func main() {
	var a, x, sum int
	x = 0
	sum = 0
	fmt.Scan(&a)

	for i := 0; i < 3; i++ {
		x = a % 10
		sum += x
		a = a / 10
	}
	fmt.Print(sum)
}
*/

/*package main

import "fmt"

func main() {
	var a, x int
	x = 0
	fmt.Scan(&a)

	for i := 0; i < 3; i++ {
		x = a % 10
		fmt.Print(x)
		a = a / 10
	}
}
*/

/*package main

import "fmt"

func main() {
	var k, h, m, t int
	h = 0
	m = 0
	fmt.Scan(&k)
	h = k / 3600
	t = k % 3600
	m = t / 60
	fmt.Print("It is ", h, " hours ", m, " minutes.")
}
*/

/*package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	c = c * c
	a = a * a
	b = b * b
	if a+b == c {
		fmt.Print("Прямоугольный")
	} else {
		fmt.Print("Непрямоугольный")
	}
}
*/

/*package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
    if a <= 0 || b <= 0 || c <= 0 {
        fmt.Println("Не существует")
    } else if a + b > c && a + c > b && b + c > a {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}
}
*/

/*package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	c := a + b
	if c%2 == 0 {
		sum := c / 2
		fmt.Print(sum)
	} else {
		sum := float32(c)
		sum = sum / 2
		fmt.Print(sum)
	}
}
*/

/*package main

import "fmt"

func main () {
	var n, a, x int
	x = 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a == 0 {
			x++
		}
	}
	fmt.Print(x)
}
*/

/*package main

import "fmt"

func bubbleSort(arr []int) {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	var n, b, c int
	var a []int
	x := 0
	fmt.Scan(&n)
	if n == 1 {
		fmt.Print("1")
	} else if n <= 0 {
		fmt.Print("0")
	} else {
		for i := 0; i < n; i++ {
			fmt.Scan(&b)
			a = append(a, b)
		}
		bubbleSort(a)
		c = a[0]
		for i := 0; i < n; i++ {
			if c == a[i] {
				x++
			}
		}
		fmt.Print(x)
	}
}
*/

/*package main

import "fmt"

func main() {
	var n, a, x, sum int
	x = 0
	fmt.Scan(&n)
	for n > 0 {
		a = n % 10
		n = n / 10
		x = x + a
	}
	a = x % 10
	x = x / 10
	sum = a + x
	fmt.Print(sum)
}
*/

/*package main

import "fmt"

func main() {
	var n, o int
	fmt.Scan(&n)
	fmt.Scan(&o)
	for i := o; i >= n; i-- {
		if i%7 == 0 {
			fmt.Print(i)
			break
		} else if i == n && i%7 != 0 {
			fmt.Print("NO")
		}
	}
}
*/
// korovy
// korova
// korov

/*package main

import "fmt"

func main() {
	var n int
	var k string
	fmt.Scan(&n)
	if n == 1 {
		k = " korova"
	} else if n >= 2 && n <= 4 {
		k = " korovy"
	} else if n >= 5 && n <= 20 {
		k = " korov"
	} else if n >= 21 {
		if n%10 == 1 {
			k = " korova"
		} else if n%10 >= 2 && n%10 <= 4 {
			k = " korovy"
		} else {
			k = " korov"
		}
	}
	fmt.Print(n, k)

}
*/

/*package main

import "fmt"

func main() {
	var n, t int
	t = 1
	fmt.Scan(&n)
	for t <= n {
		fmt.Print(t, " ")
		t = t * 2
	}
}
*/

/*package main

import "fmt"

func main() {
	var n int
	fib := []int{0, 1}
	fmt.Scan(&n)
	if n == 0 {
		fmt.Print("-1")
	}
	for i := 0; i != 40; i++ {
		a := fib[i] + fib[i+1]
		fib = append(fib, a)
	}
	for i := 0; i < 40; i++ {
		if n == fib[i] {
			fmt.Print(i)
			break
		} else if n < fib[i] {
			fmt.Print("-1")
			break
		}
	}
}
*/

/*package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	b := n
	x := 0
	for b != 0 {
		x = b % 2
		b = b / 2
		defer fmt.Print(x)
	}
}
*/

package main

import "fmt"

func main() {
	var n, o int
	fmt.Scan(&n)
	fmt.Scan(&o)
	t := n
	x := 0
	for t != 0 {
		x = t % 10
		t = t / 10
		if x == o {
			continue
		} else {
			defer fmt.Print(x)
		}
	}

}
