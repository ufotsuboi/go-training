package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"example.com/me/training-generator/generator"
)

func main() {
	const interval = 500 * time.Millisecond

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		bufio.NewScanner(os.Stdin).Scan()
		cancel()
	}()

	generator := generator.NewFizzBuzzGenerator(ctx)
	for i := range generator {
		fmt.Println(i)
		time.Sleep(interval)
	}
}
