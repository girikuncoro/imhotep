package greetings_test

import (
	"fmt"

	"github.com/derailed/imhotep/golabs/greetings"
)

func ExampleGreet() {
	fmt.Println(greetings.Greet("Fernand"))
	// Output:
	// Hello, my name is `Fernand!
}
