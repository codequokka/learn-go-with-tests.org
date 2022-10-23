package main

import "fmt"

const helloPrefix = "Hello, "

// Hello returns a personalized greeting in a given language.
func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	if language == "Spanish" {
		return "Hola, " + name
	}

	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
