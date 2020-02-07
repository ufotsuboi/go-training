package generator

import (
	"context"
	"testing"
)

func TestNewFizzBuzzGenerator(t *testing.T) {
	ctx := context.Background()
	generator := NewFizzBuzzGenerator(ctx)

	wants := [...]string{
		"1", "2", "Fizz", "4", "Buzz",
		"Fizz", "7", "8", "Fizz", "Buzz",
		"11", "Fizz", "13", "14", "FizzBuzz",
	}

	for i, want := range wants {
		if got := <-generator; got != want {
			t.Errorf("New() = generator, %v: got %v <-generator, want %v", i, got, want)
		}
	}
}
