/*
Package greetings is the worlds best GO greetings library
*/
package greetings

import "fmt"

// GreetFormat formal user greetings
const GreetFormat = "Hello, my name is `%s!"
const noName = "NoOne"

// Greet a person by name
func Greet(name string) string {
	if len(name) == 0 {
		name = noName
	}
	return format(name)
}

// CanaryGreet a person by name
func CanaryGreet(name string) string {
	if len(name) == 0 {
		name = noName
	}
	return "Hello, my name is `" + name + "!"
}

func format(name string) string {
	return fmt.Sprintf(GreetFormat, name)
}
