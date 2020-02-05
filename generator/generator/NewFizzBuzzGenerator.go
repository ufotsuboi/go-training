package generator

import (
	"context"

	"example.com/me/training-generator/fizzbuzz"
)

func NewFizzBuzzGenerator(ctx context.Context) chan string {
	ch := make(chan string)
	i := 1

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			case ch <- fizzbuzz.FizzBuzz(i):
				i++
			}
		}
	}()

	return ch
}
