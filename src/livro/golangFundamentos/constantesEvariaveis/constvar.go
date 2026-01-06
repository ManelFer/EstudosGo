// constantes e variaveis em go
package main

import (
	"fmt"
	m "math"
)

// aqui criamos a função principal
func main() {
	// criamos uma constante(constante se usa para dados que não se altera)
	const PI float64 = 3.1415 // o tipo float64 pois é um número flutuante
	var raio = 3.2

	// aqui iremos ver a forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2)
	fmt.Println("A area da circunferencia é", area)

	const (
		a = 1
		b = 2
	)

	// variaveis fica dentro da função main
	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	// forma reduzida de criar uma variavel
	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
