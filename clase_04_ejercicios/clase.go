package main

import (
	"errors"
	"fmt"
)

func main() {

	numeros := []int{}

	fmt.Println("numeros:         ", numeros)
	min, err := minimo(numeros)
	if err != nil {
		fmt.Println("A la chaucha!! algo se rompió. Qué disparate, se mató un tomate", err)
	}
	fmt.Println("minimo(numeros)err: ", err)
	fmt.Println("minimo(numeros)min: ", min)
	fmt.Println("rango(1, 10, 2): ", rango(1, 10, 2))

	x := "perro"
	y := cambiazo(&x)
	fmt.Println("x: ", x, " y: ", y)
}

func cambiazo(algo *string) string {
	*algo = "gato"
	return *algo
}

/**
 * pre-condición: la lista no debe estar vacía
 */
func minimo(lista []int) (int, error) {

	if len(lista) == 0 {
		return 0, errors.New("La lista no puede ser vacía")
	}

	menor := lista[0]

	for i := 1; i < len(lista); i++ {
		if lista[i] < menor {
			menor = lista[i]
		}
	}

	return menor, nil
}

func sum(numeros []float64) float64 {
	acumulador := 0.0
	for i := 0; i < len(numeros); i++ {
		acumulador = acumulador + numeros[i]
	}
	return acumulador
}

func rango(inicio int, fin int, step int) []int {

	resultado := make([]int, (fin-inicio)/step)
	acumulado := inicio

	for i := 0; i < len(resultado); i++ {
		resultado[i] = acumulado
		acumulado += step
	}
	return resultado
}
