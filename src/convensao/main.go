// convensao de tipos
package main

import "fmt"

var a int = 10
var b string = "Golang"

func main() {
	var c float64 = float64(a) // tipo(variavel)
	fmt.Printf("o valor de c é: %.2f e o valor de B é: %s", c, b)
}
