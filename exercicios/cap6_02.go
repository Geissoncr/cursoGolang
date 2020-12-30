package main

// escreva um programa que coloque na tela o unicode de todas as letras maiusculas do alfabeto
import "fmt"

func main() {

	for x := 65; x <= 90; x++ {
		fmt.Println(x, ":")
		for i := 0; i < 3; i++ {
			fmt.Printf("\t%#U\n", x)
		}
		fmt.Println("---------------------")
	}
}
