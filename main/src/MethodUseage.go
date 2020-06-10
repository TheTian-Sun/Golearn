package main

import (
	"fmt"
)
type Movie2 struct {
	name string
	id   int
}
func (a *Movie2) printf() {
	fmt.Println(a.name)
}

func MethodUsage() {	
	var a = new(Movie2)
	a.name = "what"
	a.id = 1

	fmt.Println(a, a.id)
	//hello()
	a.printf()

}
