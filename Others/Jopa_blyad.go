package main

import ( 
        "fmt" 
        "strings"
       )
	 
type Battery struct {
	ForTest string
}
	
func (b Battery) String() string {
	return fmt.Sprintf("%v", b.ForTest)
}

func poluch_str_zaryada () string { // получение строки с консоли
var str string
   fmt.Scan(&str)
 		   return str
}

func count (str string) int {    // подсчет количества 1 (заряда) в поданой строке
    x := strings.Count(str, "1")
    return x
}

func chargin (x int) string {     // возврат строки через массив, значения  [      XXXX]
    var zaryad []string
    pystata := 10 - x
    zaryad = append (zaryad, "[")
    for pystata != 0 {
        zaryad = append (zaryad, " ")
        pystata--
    }
    for x != 0 {
        zaryad = append (zaryad, "X")
        x--
    }
    zaryad = append (zaryad, "]")
    str := strings.Join(zaryad, "")
    return str
}

func main() {
    str := poluch_str_zaryada() // получение строки
    x := count(str) // получения числа зарядки
    s := chargin(x) // получение выводящей строки формата [      XXXX]
    batteryForTest := Battery{s}
    
    
    // batteryForTest - не забывайте об имени
// } Скобка, закрывающая функцию main() вам не видна, но она есть
