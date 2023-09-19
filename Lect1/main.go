package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args)>1 {
		fmt.Println(Say(os.Args[1]))
	} else{
		fmt.Println("Hello world!")
		//fmt.Println(os.Args[0])
	}
}