// defer: escalona as chamadas de função para serem executadas

package main

import "fmt"

func dia1() {
	fmt.Println("Domingo")
}

func dia2() {
	fmt.Println("Segunda-feira")
}

func main() {
	defer dia1()
	dia2()
}
