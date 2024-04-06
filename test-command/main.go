package main

import "fmt"

func main() {
	for i := 0; i < 10_000_000; i++ {
		fmt.Printf("Hello, world! I am test line %d of 1,000,000\n", i)
	}
}
