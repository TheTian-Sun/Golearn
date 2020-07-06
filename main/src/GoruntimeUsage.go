package main
import (
	"fmt"
	"time"
)
func GoUsage(){
	Gousage2()
	fmt.Println("This is Gouage")
	go Gousage2()
	fmt.Println("This is Gouage3")
	time.Sleep(time.Second*3)

}
func Gousage2(){
	fmt.Println("this is go usage two")
	time.Sleep(time.Second*2)
}
