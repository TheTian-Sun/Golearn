package main

import (
	"fmt"
	_ "reflect"
)

func FunctionUsage() {
	c := Add(1, 2)
	d := Add2(2, 3)
	fmt.Println(c, d)
	fmt.Println(d)
	fmt.Println(Add4(Add, 1, 9))

}

func Add(x int, y int) int {
	return x + y
}

func Add2(x int, y int) (z int) {
	z = x + y
	return
}
func Add3(numbers ...int) (int, int) {
	i := len(numbers)
	b := 0
	for _, number := range numbers {
		b += number
	}
	return i, b
}
func Add4(fn func(a int, b int) int, c int, d int) (x int) {
	x = fn(c, d)
	return
}
