// recursão: função que chama a si mesma.
// calcular fatorial de um número.

package main

import "fmt"

func fatorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * fatorial(x-1)
} // função fatorial vai recceber um número inteiro x e retornar um número inteiro
func main() {
	fmt.Println(fatorial(4))
}
