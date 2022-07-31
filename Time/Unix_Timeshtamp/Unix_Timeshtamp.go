package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	const new = 1589570165
	scanner := bufio.NewScanner(os.Stdin)                        // создание буфера для scanner
	scanner.Scan()                                               // Сканирование ст. ввода
	str_D := scanner.Text()                                      // конверт ввода в строку
	str_not_min := strings.Replace(str_D, "мин.", "m", -1)       // заменяем мин на m
	str_not_sec := strings.Replace(str_not_min, "сек.", "s", -1) // заменяем сек на s
	str := strings.Replace(str_not_sec, " ", "", -1)             // удаляем пробелы
	dur, _ := time.ParseDuration(str)                            // конвертируем строку в duration
	unix_itog := new + int64(dur.Seconds())                      // узнаем новый UNIX SHTAMP
	tim_itog := time.Unix(unix_itog, 0)                          // переводим в ст время UNIX
	fmt.Print(tim_itog.Format(time.UnixDate))                    // выводим по формату UnixData
}
