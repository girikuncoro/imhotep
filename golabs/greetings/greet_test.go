package greetings_test

import (
	"fmt"
	"testing"

	"github.com/derailed/imhotep/golabs/greetings"
)

var useCases = []struct {
	name, expected string
}{
	{"Fernand", "Fernand"},
	{"", "NoOne"},
}

func TestGreet(t *testing.T) {
	for _, uc := range useCases {
		expected := fmt.Sprintf(greetings.GreetFormat, uc.expected)
		actual := greetings.Greet(uc.name)
		if actual != expected {
			t.Fatalf("Expecting `%s` GOT `%s`", expected, actual)
		}
	}
}

func TestCanaryGreet(t *testing.T) {
	for _, uc := range useCases {
		expected := fmt.Sprintf(greetings.GreetFormat, uc.expected)
		actual := greetings.CanaryGreet(uc.name)
		if actual != expected {
			t.Fatalf("Expecting `%s` GOT `%s`", expected, actual)
		}
	}
}

func BenchmarkGreet(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, uc := range useCases {
			expected := fmt.Sprintf(greetings.GreetFormat, uc.expected)
			actual := greetings.Greet(uc.name)
			if actual != expected {
				b.Fatalf("Expecting `%s` GOT `%s`", expected, actual)
			}
		}
	}
}

func BenchmarkCanaryGreet(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, uc := range useCases {
			expected := fmt.Sprintf(greetings.GreetFormat, uc.expected)
			actual := greetings.CanaryGreet(uc.name)
			if actual != expected {
				b.Fatalf("Expecting `%s` GOT `%s`", expected, actual)
			}
		}
	}
}
