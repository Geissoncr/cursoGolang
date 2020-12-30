package main

// escreva um programa que crie um loop utilizando a sintax for condition{}
// Utilizeo para demonstrar os anos que voce nasceu até hoje
import "fmt"

func main() {
	anoNascimento := 1980
	anoFinal := 2020
	for anoNascimento <= anoFinal {
		fmt.Println(anoNascimento)
		anoNascimento++
	}
}
