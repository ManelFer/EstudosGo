// estudos de arrays em Go
package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 5.3
	x[1] = 6.2
	x[2] = 7.1
	x[3] = 8.0
	x[4] = 9.4

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]

	}
	fmt.Println("a media é:", total/float64(len(x)))
}
