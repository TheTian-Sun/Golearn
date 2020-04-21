package main

import (
	"fmt"
	_ "reflect"
	"strconv"
)

func TypeConvert() {
	s := "true"
	b := true
	ss, _ := strconv.ParseBool(s)
	fmt.Println(b == ss)
	s = strconv.FormatBool(b)
	fmt.Println(s)
	fmt.Println(b)

}
