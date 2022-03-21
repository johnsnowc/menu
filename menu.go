package main

import "fmt"

func main() {
	var str string
	fmt.Println("Menu Program")
	for {
		fmt.Scan(&str)
		switch str {
		case "hello":
			fmt.Println("hello")
		case "world":
			fmt.Println("world")
		case "quit":
			fmt.Print("ok,goodbye")
			return
		default:
			fmt.Println("wrong choice!")
		}
	}
}
