package main
import(
	"fmt"
)
 func MapUsage(){
	 var players = make(map[string]int)
	 players["Leo"]=1
	 players["Tian"]=0
	 fmt.Println(players)
	 delete(players,"Leo")
	 fmt.Println(players)
 }
