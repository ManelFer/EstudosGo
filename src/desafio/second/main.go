// criar um programa na linguagem Go que trabalhe com operadir % e for
// objetivo: precisa exibir todos os números entre 1 e 100 que são divisíveis por 3

package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
