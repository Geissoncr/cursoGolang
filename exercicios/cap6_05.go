package main

// escreva um programa que demonstre o resto da divisão por 4 de todos os numeros entre 10 e 100
import "fmt"

func main() {

	for x := 10; x <= 100; x++ {
		fmt.Println(x, " ", x%4)
	}
}
