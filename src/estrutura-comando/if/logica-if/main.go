package main

import "fmt"

func main() {
	x := 9
	if x == 2 || x == 3 {
		fmt.Println("X é 2 ou 3")
	} else {
		fmt.Println("X não é 2 nem 3")
	}

	if x%3 == 0 && x%2 == 0 {
		fmt.Println("X é múltiplo de 2 e 3")
	} else if x%2 == 0 {
		fmt.Println("X é múltiplo de 2 e não é múltiplo de 3")
	} else if x%3 == 0 {
		fmt.Println("X é múltiplo de 3 e não é múltiplo de 2")
	} else {
		fmt.Println("X não é múltiplo de 2 nem de 3")
	}
}
