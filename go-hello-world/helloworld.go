package main

import "fmt"

func main() {
	fmt.Println(Hello("Andres"))
}

func Hello(name string) string {
	return "Hello, " + name
}
