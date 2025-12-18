// estudos sobre if em go
package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d: ", i)
		if i%2 == 0 {
			fmt.Println("Número par")
		} else {
			fmt.Println("Número ímpar")
		}

	}
}
