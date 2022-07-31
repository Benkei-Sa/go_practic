package main

import (
    "fmt"
    "time"
    )

func out(from, to int, ch chan bool) {
    for i:=from; i<=to; i++ {
        time.Sleep(50 * time.Millisecond)
        fmt.Println(i)
    }
    ch <- true
}

func main() {
    ch := make(chan bool)

    go out(0, 5, ch)
    go out(6, 10, ch)

    <-ch
    <-ch
}

//

func evenSum(from, to int, ch chan int) {
	result := 0
	for i:=from; i<=to; i++ {
	  if i%2 == 0 {
		result += i
	  }    
	}
	ch <- result
  }
  func squareSum(from, to int, ch chan int) {
	result := 0
	for i:=from; i<=to; i++ {
	  if i%2 == 0 {
		result += i*i
	  }    
	}
	ch <- result
  } 

  //

evenCh := make(chan int)
sqCh := make(chan int)

go evenSum(0, 100, evenCh)
go squareSum(0, 100, sqCh)

fmt.Println(<-evenCh + <- sqCh)

//

evenCh := make(chan int)
sqCh := make(chan int)

go evenSum(0, 100, evenCh)
go squareSum(0, 100, sqCh)

select {
  case x := <- evenCh:
     fmt.Println(x)
  case y := <- sqCh:
     fmt.Println(y)
}

//

for {
    select {
        case x := <- evenCh:
            fmt.Println(x)
            return
        case y := <- sqCh:
            fmt.Println(y)
            return
        default:
            fmt.Println("Ничего не доступно")
            time.Sleep(50 * time.Millisecond)
    }
}




// Последняя задача 



package main
import "fmt"

//определите функцию download()
func download(s int, ch chan int){
    sum := 0
    for i:=s; i > 0; i-- {
        sum += i
    }
    ch <- sum
}


func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    ch3 := make(chan int)

    var s1, s2, s3 int
    fmt.Scanln(&s1)
    fmt.Scanln(&s2)
    fmt.Scanln(&s3)

    go download(s1, ch1)
    go download(s2, ch2)
    go download(s3, ch3)

    //выведите сумму всех результатов
    sum := <-ch1 + <-ch2 + <-ch3
    fmt.Println(sum)
}