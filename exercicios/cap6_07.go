package main

// escreva um programa que demonstre o funcionamento da função IF
// agora tbm use else if e else
import "fmt"

func main() {

	nome := "Geisson!"

	if nome == "Geisson" {
		fmt.Println("Nome é Geisson!")
	} else if nome == "Carlos" {
		fmt.Println("Nome não é Geisson!")

	} else {
		fmt.Println("Nenhum nome")
	}
}
