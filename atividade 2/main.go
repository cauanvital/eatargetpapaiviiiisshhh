package main

import "fmt"

func isFibonacci(n int) bool {

	if n < 0 {
		return false
	}
	if n == 0 || n == 1 {
		return true
	}

	a, b := 0, 1

	for {

		next := a + b

		if next == n {
			return true
		}
		if next > n {
			return false
		}

		a = b
		b = next

	}

}

func main() {

	n := 13

	ok := isFibonacci(n)
	if ok {
		fmt.Println("Is Fibonacci!")
	} else {
		fmt.Println("Oh no...")
	}

}
