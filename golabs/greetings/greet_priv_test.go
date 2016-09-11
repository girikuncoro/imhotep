package greetings

import (
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {
	var (
		name     = "Fernand"
		expected = fmt.Sprintf(GreetFormat, name)
		actual   = format(name)
	)

	if actual != expected {
		t.Fatalf("Expecting `%s` GOT `%s`", expected, actual)
	}
}
