// coleções de campos
// agrupar dados
// formar registros
//
//	type NomeDaEstrutura struct {
//		campo1 tipoDoCampo1
//		campo2 tipoDoCampo2
//		campo3 tipoDoCampo3
//	}
package main

import "fmt"

type Pessoa struct {
	nome  string
	idade int
}

func main() {
	fmt.Println(Pessoa{"João", 30})
	fmt.Println(Pessoa{nome: "Maria", idade: 25})
}
