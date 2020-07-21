package main

import (
	"fmt"
	"time"
)

func GoUsage() {
	Gousage2()
	fmt.Println("This is Gouage")
	go Gousage2()
	c := make(chan string)
	fmt.Println("This is Gouage3")
	time.Sleep(time.Second * 3)
	go Gousage3(c)
	msg := <-c
	fmt.Println(msg)

}
func Gousage2() {
	fmt.Println("this is go usage two")
	time.Sleep(time.Second * 2)
}
func Gousage3(c chan string) {
	time.Sleep(time.Second * 2)
	fmt.Println("This is GoUsage Three")
	c <- "GoUsageThree End"
}
