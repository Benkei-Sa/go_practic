package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, name := creation_file() // Создание файла с именем введеным пользователем
	scan_and_write_num(file, name)
}

func creation_file() (*os.File, string) {
	fmt.Print("Введите название файла: ")
	var name string
	fmt.Scan(&name)
	file, err := os.Create(name)
	if err != nil {
		fmt.Errorf("Не удалось создать файл: %v", err)
	}
	return file, name
}

func scan_and_write_num(file *os.File, name string) {
	f, err := os.Open(name)
	if err != nil {
		fmt.Errorf("Не удалось открыть файл из дерриктории: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(os.Stdin)
	//	fmt.Print("Введите числа для сумирования ")
	for scanner.Scan() {
		fmt.Println("Вопрос:")
		fmt.Println(scanner.Text())
		f.WriteString(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Errorf("error detected: %v", err)
		return
	}
	//scanner.Scan()
	//f.WriteString(scanner.Text())

}

func read_file() {

	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Errorf("Не удалось прочитать строку из os.Stdin: %v", err)
	} else if err == io.EOF { // если файл закончился -> закончить
		break
	}

}
