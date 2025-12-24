// closure: função dentro de outra função que tem acesso às variáveis da função externa.

package main

import "fmt"

func main() {
	x := 0

	numero := func() int { // precisa ter o tipo de retorno
		x++ // incremento somar 1
		return x
	}

	fmt.Println(numero())
}
