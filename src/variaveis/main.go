package main

import (
	"fmt"
)

// declarar uma variavel que não se altera
const ebulicaoF float64 = 212.0

// declaração de variáveis globais
func main() {
	temperaturaF := ebulicaoF
	temperaturaC := (temperaturaF - 32) * 5 / 9

	fmt.Printf("A temperatura de ebulição da água é %.2f°F ou %.2f°C\n", temperaturaF, temperaturaC)
}
