package main

import "fmt"

func main() {

	nome := "Geisson Silva"
	matricula := 11320000
	var idade int

	fmt.Println("Olá sr. ", nome, "sua matricula é ", matricula)

	fmt.Println("Qual a Sua Idade? ")
	fmt.Scanf("%d", &idade)
	fmt.Println("Sua Idade é ", idade)
}
