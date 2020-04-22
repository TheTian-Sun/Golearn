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
	fmt.Println(m)
} 