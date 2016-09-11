package fizzbuzz_test

import (
	"testing"

	"github.com/derailed/imhotep/golabs/fizzbuzz"
)

var useCases = []struct {
	number   int
	expected string
}{
	{1, "1"},
	{3, "Fizz"},
	{5, "Buzz"},
	{15, "FizzBuzz"},
}

func TestFizzbuzz(t *testing.T) {
	for _, uc := range useCases {
		if actual := fizzbuzz.Play(uc.number); actual != uc.expected {
			t.Fatalf("(%d) Expected %s go %s", uc.number, uc.expected, actual)
		}
	}
}
