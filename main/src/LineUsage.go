package main

import (
	"fmt"
	"os"
)

func LineUsage() {
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
}
