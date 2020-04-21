package main

import (
	"fmt"
	"reflect"
)

func slice() {
	var a = [5]int{1, 3}
	//a[1] = "hello"
	//a[2] = "man"
	b := a[1:3]
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(cap(b), len(b))
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
}
