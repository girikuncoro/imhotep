package greetings

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	var (
		expected = fmt.Sprintf(greetFormat, "Fernand")
		actual   = greet("Fernand")
	)

	if actual != expected {
		t.Fatalf("Expecting `%s` GOT `%s`", expected, actual)
	}
}
