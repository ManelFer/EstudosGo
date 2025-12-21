// criar um programa na linguagem Go que trabalhe com operadir % e for
// objetivo: precisa exibir todos os números entre 1 e 100 que são divisíveis por 3

package main

import (
	"fmt"
)

/*func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}*/

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("pinpan")
		} else if i%3 == 0 {
			fmt.Println("pin")
		} else if i%5 == 0 {
			fmt.Println("pan")
		} else {
			fmt.Println(i)
		}
	}
}

// segunda parte do desafio
// precisa exibir de 1 a 100 e os multiplos de 3 aparecer pin e os multiplos de 5 pan e se for ambom 3 e 5 pinpan
