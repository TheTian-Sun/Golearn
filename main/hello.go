package main

import (
	"fmt"
	"os"
)

func hello() {
	var b bool 
	b = true
	var s string
	s = "hello"
	fmt.Println(b)
	fmt.Println(s)
	a := 23
	fmt.Println(a)
	var se, sep string
	for i := 1; i< len(os.Args); i++ {
		se += sep
		sep = " "
	}
	fmt.Println(se)
}


