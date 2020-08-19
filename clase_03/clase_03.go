package main

import "fmt"

func main() {
	arr := [8]float64{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(reverse(arr[:]))
}

func reverse(numeros []float64) []float64 {
	largo := len(numeros)
	resultado := make([]float64, largo)

	for i := 0; i < largo; i++ {
		resultado[largo-i-1] = numeros[i]
	}

	return resultado
}

func sum(numeros []float64) float64 {
	acumulador := 0.0
	for i := 0; i < len(numeros); i++ {
		acumulador = acumulador + numeros[i]
	}
	return acumulador
}

func avg(numeros []float64) float64 {
	return sum(numeros) / float64(len(numeros))
}
