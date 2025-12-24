// fun também pode ser chamada de procedimento ou subrotina
// parte do codigo
// ela pega um dado de entrada, processa e devolve um dado de saída

package main

import "fmt"

func media(lista []float64) float64 {
	total := 0.0
	for _, valor := range lista {
		total += valor
	}
	return total / float64(len(lista))
}

func main() { // programa de media de sala
	lista := []float64{7.5, 8.0, 6.5, 9.0, 5.5} //lista
	fmt.Println(media(lista))

	//total := 0.0
	//for _, valor := range lista {
	//	total += valor
	//}

	//fmt.Println(media(lista)) // imprimir media da sala
}
