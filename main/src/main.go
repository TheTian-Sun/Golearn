package main

import (
	"fmt"

	"github.com/golang/example/stringutil"
)

func main() {
	defer fmt.Println("hello world")
	array()
	slice()
	TypeConvert()
	FunctionUsage()
	MapUsage()
	defer fmt.Println("你好")
	fmt.Println(stringutil.Reverse("nihao"))
	StructUsage()
	MethodUsage()
	InterfaceUsage()
	GoUsage()
}
