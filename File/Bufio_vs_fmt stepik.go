package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {

	scan_and_write_num()

}

func scan_and_write_num() {
	var sum int
	var result string
	file, _ := os.Create("Name.txt")
	f, _ := os.Open("Name.txt")
	defer f.Close()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		x, _ := strconv.Atoi(str)
		sum = sum + x
		result = strconv.Itoa(sum)
		if scanner.Scan() != true {
			break
		}
	}
	io.WriteString(os.Stdout, result)
}

/*func read_file(file *os.File) {

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { // если файл закончился -> закончить
			break
		}
		x, _ := strconv.Atoi(str)
		sum = sum + x
	}
	result := strconv.Itoa(sum)
	io.WriteString(os.Stdout, result)
}
*/
