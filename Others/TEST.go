/*arr := make([]int, 10)
fin := make(map[int]int)

for i := range arr {
   fmt.Scan(&arr[i])
}
for i := range arr {
	_, ok := fin[arr[i]]
    if ok == false {
        value := work(arr[i])
    	fin[arr[i]] = value
    }
	fmt.Print(fin[arr[i]], " ")
}


for range cityPopulation {
	for j := 0; j < len(groupCity[10]); j++ {
		_, ok :=  cityPopulation[groupCity[10][j]]
		if ok == true {
			delete(cityPopulation, groupCity[10][j])
		}
	}
	for j := 0; j < len(groupCity[1000]); j++ {
		_, ok :=  cityPopulation[groupCity[1000][j]]
		if ok == true {
			delete(cityPopulation, groupCity[1000][j])
		}
	}
}
*/

//package main

//import (
//	"fmt"
//	"strconv"
//)

/*func main() {
	res := []rune("stepik")
	fmt.Print(res)
}
*/
/*func droop (a, b []byte) []byte {
	for i := range a {
		if a[i] >= 48 && a[i] <= 57 {
			b = append(b, a[i])
		}
	}
	return b
}
func conv_type_byte_to_s (a []byte) int64 {
	s_num, err := strconv.ParseInt(string(a), 10, 64)
	if err != nil {
		panic(err)
	}
	return s_num
}

func adding (one, two string) int64 {
	var num_o, num_t []byte
	fmt.Scan(&one, &two)
	b_one := []byte(one)
	b_two := []byte(two)
	num_o = droop(b_one, num_o)
	num_t = droop(b_two, num_t)
	s_num_o := conv_type_byte_to_s(num_o)
	s_num_t := conv_type_byte_to_s(num_t)
	fmt.Print(s_num_o + s_num_t)
}

func adding (one, two string) int64 {
	//func main() {
		//var one, two string
		var num_o, num_t []byte
		fmt.Scan(&one, &two)
		b_one := []byte(one)
		b_two := []byte(two)
		for i := range b_one {
			if b_one[i] >= 48 && b_one[i] <= 57 {
				num_o = append(num_o, b_one[i])
			}
		}
		for i := range b_two {
			if b_two[i] >= 48 && b_two[i] <= 57 {
				num_t = append(num_t, b_two[i])
			}
		}
		s_num_o, err := strconv.ParseInt(string(num_o), 10, 64)
		if err != nil {
			panic(err)
		}
		s_num_t, err := strconv.ParseInt(string(num_t), 10, 64)
		return (s_num_o + s_num_t)
	}

*/

/*package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var text string
	text, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	result(text)
}

func result(text string) {
	str := strings.Replace(text, " ", "", -1)
	m_str := strings.Split(str, ";")
	one := rep_and_conv(m_str, 0)
	two := rep_and_conv(m_str, 1)
	fmt.Printf("%.4f", (one / two))
}

func rep_and_conv(m_str []string, i int) float64 {
	a := m_str[i]
	num_one := []rune(a)
	for i := range num_one {
		if num_one[i] == ',' {
			num_one[i] = '.'
		}
	}
	str_one := string(num_one)
	str_two := droop(str_one)
	num, err := strconv.ParseFloat(str_two, 64)
	if err != nil {
		panic(err)
	}
	return num
}

func droop(str string) string {
	var str_b_two []byte
	str_byte := []byte(str)
	for i := range str_byte {
		if ((str_byte[i] >= 48) && (str_byte[i] <= 57)) || (str_byte[i] == 46) {
			str_b_two = append(str_b_two, (str_byte[i]))
		}
	}
	str_two := string(str_b_two)

	return str_two
}
*/

package main

import "fmt"

func main() {
	x := func(fn func(i int) int, i int) func(int) int { return fn }(func(i int) int { return i + 1 }, 5)
	fmt.Printf("%T", x)
}


fn := func(a uint) uint {
    var a uint
	if a == 0 {
		res_fin = 100
	} else {
    	temp := uint64(a)
   		var result []byte
  		str := strconv.FormatUint(temp, 36)
   		str_byt := []byte(str)
    	for i := range str_byt {
        	for j = 50; j <= 56; _ {
				if str_byte[i] == j {
            		result = append (result, str_byt[i])
        		}
				j = j + 2
			}
    	}
		str = string(str_byt)
		res_fin, _ = strconv.ParseUint(str, 10, 0)
       	return (res_fin)
	}
}



fn := func(x uint) uint {
    var result, o_del, step uint
    step = 1
    var itog []uint
    i := x
    if i == 0 {
        result = 100
    } else {
        for i != 0 {
            o_del = i % 10
            if o_del % 2 == 0 && o_del != 0 {
                itog = append (itog, o_del)
            }
            i = i / 10
        }
        if len(itog) < 1 {
            result = 100
        } else {
            for i := 0; i < len(itog); i++ {
                if i == 0 {
                    result = itog[i]
                } else {
                    step = 1
                    for j := i; j >= 1; j-- {
                        step = step * 10
                    }
                    result += (itog[i] * step)
                }
            }
        }
    }
    return result
}



