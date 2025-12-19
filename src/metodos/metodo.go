//método é uma função associada a um tipo particular
// Isto é, em Poo objeto é um valor e método é uma função associada a esse valor

package main

import (
	"fmt"
)

type retangulo struct {
	comprimento, altura int
}

func (r *retangulo) area() int {
	return r.comprimento * r.altura
}

func (r retangulo) perimetro() int {
	return 2*r.comprimento + 2*r.altura
}

func main() {
	r := retangulo{comprimento: 10, altura: 5}

	fmt.Println("Área do retângulo:", r.area())
	fmt.Println("Perímetro do retângulo:", r.perimetro())
}

// assosiamos a variavel r do tipo retangulo dois métodos: area e perimetro
// o método area tem um receiver do tipo ponteiro *retangulo, ou seja, ele recebe o endereço de memória do valor
// o método perimetro tem um receiver do tipo valor retangulo, ou seja, ele recebe uma cópia do valor
// ambos os métodos podem ser chamados a partir de uma variável do tipo retangulo, como no exemplo acima
