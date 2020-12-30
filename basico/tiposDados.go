// O exemplo mostra vários tipos de variáveis e também que as declarações de variáveis
// podem ser "construídas" em blocos, como com as declarações de importação.

// Os tipos int, uint e uintptr são geralmente de 32 bits em sistemas de 32 bits e 64 bits
// em sistemas de 64 bits. Quando você precisar de um valor inteiro deverá usar int, a menos
// que tenha um motivo específico para usar um tipo de inteiro com tamanho especificado ou sem sinal.

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
