package main

import (
	"fmt"
	"time"
)

func main() {

	for index := 0; index < 10; index++ {
		fmt.Println("hello world")
		fmt.Println("hi")
	}
	time.Sleep(time.Second * 20)
}
