package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome = "Geisson Silva"
	var matricula = 122336655
	var versao = 1.1
	fmt.Println("Nome do Aluno ", nome, " e sua matricula é ", matricula)
	fmt.Println("Este programa é a versão ", versao)
	fmt.Println("A raiz quadrada  ", reflect.TypeOf(versao))

}
