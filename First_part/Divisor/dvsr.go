package main

import "fmt"

// this is a function finding the number of divisor
func Divisors(n int) int {
	var counter int = 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			counter = counter + 1
		}

	}
	return counter
}

func main() {

	x := 1
	y := Divisors(x)
	fmt.Println(y)
}
