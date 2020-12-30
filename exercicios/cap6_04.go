package main

// escreva um programa que crie um loop utilizando a sintax for{}
// Utilizeo para demonstrar os anos que voce nasceu até hoje
import "fmt"

func main() {
	anoNascimento := 1980
	anoFinal := 2020
	for {
		if anoNascimento > anoFinal {
			break
		}
		fmt.Println(anoNascimento)
		anoNascimento++
	}
}
