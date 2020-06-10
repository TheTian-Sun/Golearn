package main

import (
	"fmt"
)

type T850 struct {
	Name string
}
type Robot interface {
	PowerOn() error
}

func (a *T850) PowerOn() error {
	return nil
}
func InterfaceUsage() {
	a := T850{
		Name: "12a",
	}
	Err := a.PowerOn()
	fmt.Println(Err)
	fmt.Println(a)
}
