/*
Package greetings is the worlds best GO greetings library
*/
package greetings

import "fmt"

const greetFormat = "Hello, my name is `%s!"
const noName = "NoOne"

// greet a person by name
func greet(name string) string {
	if len(name) == 0 {
		name = noName
	}
	return format(name)
}

func format(name string) string {
	return fmt.Sprintf(greetFormat, name)
}
