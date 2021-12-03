package main

import (
	"strconv"

	"github.com/suborbital/reactr/api/tinygo/runnable"
)

type Tinygo struct{}

func (h Tinygo) Run(input []byte) ([]byte, error) {
	inStr := string(input)
	n, _ := strconv.Atoi(inStr)
	f := factorial(n)

	return []byte("factorial of, " + inStr + " is " + string(f)), nil
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}

	return n * factorial(n-1)
}

// initialize runnable, do not edit //
func main() {
	runnable.Use(Tinygo{})
}
