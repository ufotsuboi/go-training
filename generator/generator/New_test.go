package generator

import (
	"context"
	"testing"
)

func TestNew(t *testing.T) {
	ctx := context.Background()
	generator := New(ctx)

	wants := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, want := range wants {
		if got := <-generator; got != want {
			t.Errorf("New() = generator, %v: got %v <-generator, want %v", i, got, want)
		}
	}
}
