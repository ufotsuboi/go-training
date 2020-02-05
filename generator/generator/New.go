package generator

import (
	"context"
)

func New(ctx context.Context) chan int {
	ch := make(chan int)
	i := 1

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			case ch <- i:
				i++
			}
		}
	}()

	return ch
}
