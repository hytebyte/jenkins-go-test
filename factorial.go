package main

import "fmt"

func factorial(n int) int {
	if n < 0 {
		panic(fmt.Sprintf("n = %d passed to factorial", n))
	}
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
