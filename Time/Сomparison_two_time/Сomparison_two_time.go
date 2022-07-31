package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func scaner_time() string {
	scanner := bufio.NewScanner(os.Stdin) // создание буфера для scanner
	scanner.Scan()                        // получение строки формата *bufio.Scanner
	d_time := scanner.Text()              // перевод *bufio.Scanner в string
	return d_time
}

func str_split_t_conv_time(d_time string) (time.Time, time.Time) {
	time_mas := strings.Split(d_time, ",")
	tt_one := conv_str_t_time(time_mas[0])
	tt_two := conv_str_t_time(time_mas[1])
	return tt_one, tt_two
}

func conv_str_t_time(time_zn string) time.Time {
	time, err := time.Parse("02.01.2006 15:04:05", time_zn)
	if err != nil {
		fmt.Print(err)
	}
	return time
}

func result_compres(tt_one, tt_two time.Time, compres bool) {
	if compres == true {
		fmt.Println(tt_one.Sub(tt_two))
	} else {
		fmt.Println(tt_two.Sub(tt_one))
	}
}

func main() {
	d_time := scaner_time()                         //получаем строковое время со стандартного ввода
	tt_one, tt_two := str_split_t_conv_time(d_time) //получаем из строки 2 значения формата time.Time
	compres := tt_one.After(tt_two)                 //проверяем какое время из 2-х больше
	result_compres(tt_one, tt_two, compres)
}
