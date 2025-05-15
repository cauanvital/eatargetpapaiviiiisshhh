package main

import "fmt"

func main() {
	i, soma, k := 13, 0, 0
	for k < i {
		k = k + 1
		soma = soma + k
	}
	fmt.Println(soma)
}
