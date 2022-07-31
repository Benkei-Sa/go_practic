/*package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	runa := []rune(text)
	if (unicode.IsUpper(runa[0]) == true) && (strings.HasSuffix(text, ".") == true) {
		fmt.Print("Right")
	} else {
		fmt.Print("Wrong")
	}
}
*/

/*package main

import (
	"fmt"
)
func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main () {
	var text, revtext string
	fmt.Scan(&text)
	revtext = reverse(text)
	if text == revtext {
		fmt.Print("Палиндром")
	} else {
		fmt.Print("Нет")
	}
}
*/

/*package main

import (
	"fmt"
	"strings"
)

func main() {
	var X, S string
	fmt.Scan(&X)
	fmt.Scan(&S)
	fmt.Print(strings.Index(X, S))
}
*/

/*package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	for i, c := range str {
		if i%2 == 1 {
			fmt.Printf("%c", c)
		}
	}
}
*/

/*package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	//	fmt.Printf("%c", x)
	for i, c := range str {
		if strings.Count(str, string(c)) == 1 {
			x := str[i]
			fmt.Printf("%c", x)
		}
	}
}
*/

package main

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

func main() {
	var pas string
	fmt.Scan(&pas)
	l := utf8.RuneCountInString(pas)
	fmt.Println(pas)
	matched, _ := regexp.MatchString(`^\w{5,}$`, pas)
	fmt.Println(matched)
	if l >= 5 && matched == true {
		fmt.Print("Ok")
	} else {
		fmt.Print("Wrong password")
	}
}
