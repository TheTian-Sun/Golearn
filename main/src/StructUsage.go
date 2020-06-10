package main
import (
	"fmt"
)

func StructUsage() {
	type Movie struct {
		name string
		id int
	}
	var m Movie
	m.id =0001
	m.name="龙虎门"
	b:= m
	fmt.Println(m)
	fmt.Println(b)
	c := &m
	fmt.Println(*c)
	c.id = 0
	fmt.Println(&m)
	fmt.Println(c)
} 