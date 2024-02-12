package main

import "fmt"

const (
	french  = "French"
	spanish = "Spanish"

	prefixHello   = "Hello, "
	prefixHola    = "Hola, "
	prefixBonjour = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = prefixBonjour
	case spanish:
		prefix = prefixHola
	default:
		prefix = prefixHello
	}
	return
}

func main() {
	fmt.Println(Hello("Andres", ""))
}
