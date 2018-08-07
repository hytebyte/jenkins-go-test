package main

import (
	"fmt"
	"io"
)

func main() {
	for {
		fmt.Print("n = ")
		var n int
		if _, err := fmt.Scanf("%d", &n); err == io.EOF {
			break
		}
		fmt.Println("n! =", factorial(n))
	}
	fmt.Println("done!")
}
