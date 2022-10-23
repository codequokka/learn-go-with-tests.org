package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello returns a personalized greeting in a given language.
func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	greetingPrefix := greetingPrefix(language)

	return greetingPrefix + name
}

// greetingPrefix returns a greeting prefix in a given language.
func greetingPrefix(language string) (greetingPrefix string) {
	switch language {
	case spanish:
		greetingPrefix = spanishHelloPrefix
	case french:
		greetingPrefix = frenchHelloPrefix
	default:
		greetingPrefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
