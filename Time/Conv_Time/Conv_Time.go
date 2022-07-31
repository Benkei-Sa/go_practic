package main

import (
	"fmt"
	"time"
)

func scan() string {
	var s_time string
	fmt.Scan(&s_time)
	return s_time
}

func conv_str_to_time(s_time string) time.Time {
	tim, err := time.Parse(time.RFC3339, s_time)
	if err != nil {
		fmt.Print(err)
	}
	return tim
}

func main() {
	s_time := scan()
	tim := conv_str_to_time(s_time)
	conv_Time := tim.Format(time.UnixDate)
	fmt.Printf("%s", conv_Time)
}
