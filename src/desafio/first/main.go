// o desafio é desenvolver um porgrama para converter o valor do ponto de ebulição da água de Kelvin para Celsius

package main

import "fmt"

const ebulicaoK float64 = 373.15

func main() {
	temperaturaK := ebulicaoK
	temperaturaC := temperaturaK - 273.15
	fmt.Printf("A temperatura de ebulição da água é %.2f ºk ou %.2f ºC\n", temperaturaK, temperaturaC)
}
