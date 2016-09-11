package greetings

import (
	"fmt"
	"testing"
)

var useCases = []struct {
	name, expected string
}{
	{"Fernand", "Fernand"},
	{"", "NoOne"},
}

func TestGreet(t *testing.T) {
	for _, uc := range useCases {
		expected := fmt.Sprintf(greetFormat, uc.expected)
		actual := greet(uc.name)
		if actual != expected {
			t.Fatalf("Expecting `%s` GOT `%s`", expected, actual)
		}
	}
}
