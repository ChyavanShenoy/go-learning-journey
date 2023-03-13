package main

import "fmt"

func main() {
	name := ""
	fmt.Println("Hi there! What's your name?")
	fmt.Scanln(&name)
	message := fmt.Sprintf("Hello, %s! Hope you enjoy the journey of learning Go.", name)
	fmt.Println(message)
}
