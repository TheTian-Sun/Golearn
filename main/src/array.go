package main

import (
	"fmt"
)

//数组相关
func array() {
	//var a[3] int //数组声明

	var city = [4]string{"北京", "上海", "广州", "深圳"}
	fmt.Println(city)
	var city1 = [...]string{"北京", "上海", "广州", "深圳"}
	fmt.Println(city1)
	var city2 = [...]string{1: "北京", 4: "上海", 7: "广州", 8: "深圳"}
	fmt.Println(city2)

	fmt.Printf("%T,%T,%T", city, city1, city2)

	for i := 0; i < len(city); i++ {
		fmt.Println(city[i])
	}
	for index, value := range city {
		fmt.Println(index, value)
	}
   var data = [4]int{1,3,5,8}
    number := 0
   for _,value := range data{
	number= number + value
	fmt.Println(number)
   }
}

func main() {
	array()
	
}
