package main

import "fmt"

func hello() {
	var hello string = "Hello world."
	for i := 0; i < 10; i++ {
		fmt.Println(hello, i)
	}
}

func main() {
	hello()
}
