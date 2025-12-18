// não existe with em Go, mas podemos simular seu comportamento usando for

package main

import "fmt"

func main() {
	i := 0
	for i < 20 {
		fmt.Println("i é menor que 20")
		i++
	}
}
