package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func conv_str_to_tim(d_time string) time.Time {
	tim, err := time.Parse("2006-01-02 15:04:05", d_time)
	if err != nil {
		fmt.Print(err)
	}
	return tim
}

func exam_to_Print(tim time.Time) {
	if tim.Hour() > 13 {
		new_date_tim := tim.AddDate(0, 0, 1)
		fmt.Print(new_date_tim.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Print(tim.Format("2006-01-02 15:04:05"))
	}
}

func scanner_time() string {
	scanner := bufio.NewScanner(os.Stdin) // создание буфера для scanner
	scanner.Scan()                        // получение строки формата *bufio.Scanner
	d_time := scanner.Text()              // перевод *bufio.Scanner в string
	return d_time
}

func main() {
	d_time := scanner_time()       // перевод *bufio.Scanner в string
	tim := conv_str_to_tim(d_time) // конвертирование string в time.Time
	exam_to_Print(tim)             // проверяем время по условию и выводим результат
}
