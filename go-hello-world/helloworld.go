package main

import "fmt"

const prefixHello = "Hello, "
const prefixHola = "Hola, "
const prefixBonjour = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	if language == "Spanish" {
		return prefixHola + name
	}
	if language == "French" {
		return prefixBonjour + name
	}
	return prefixHello + name
}

func main() {
	fmt.Println(Hello("Andres", ""))
}
