/*
Package greetings is the worlds best GO greetings library
*/
package greetings

import "fmt"

const greetFormat = "Hello, my name is `%s!"

// greet a person by name
func greet(name string) string {
	return fmt.Sprintf(greetFormat, name)
}
