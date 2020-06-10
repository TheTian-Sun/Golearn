package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("hello world")
	array()
	slice()
	TypeConvert()
	FunctionUsage()
	MapUsage()
	defer fmt.Println("你好")
	StructUsage()
	MethodUsage()
	InterfaceUsage()
}
