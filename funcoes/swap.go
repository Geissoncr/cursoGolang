package main

import "fmt"

func swap(x, y string) (string, string) {

	return y, x
}

func main() {
	a, b := "teste", "legal"

	fmt.Println(swap(a, b))
}
