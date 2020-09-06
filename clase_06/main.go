package main

import (
	"fmt"
	"strings"
)

func main() {

	numeros := []int{1, 2, 3, 4, 5, 4, 2}

	fmt.Println(sumar(numeros))

	fmt.Println(reduce(numeros, func(accum int, item int) int {
		return accum + item
	}))

	accum := 0
	forEach(numeros, func(index int, item int) {
		accum += item
	})

	fmt.Println(duplicar(numeros))

	listaDoble := make([]int, len(numeros))
	forEach(numeros, func(index int, item int) {
		listaDoble[index] = item * 2
	})
	fmt.Println(listaDoble)

	duplicados := make([]string, len(numeros))
	forEachIndexed(numeros, func(index int, item int) {
		duplicados[index] = strings.Repeat(".", item)
	})
	fmt.Println(duplicados)
}

func forEachIndexed(items []int, f func(int, int)) {
	//acumulador := make([]int, len(items))

	for i := 0; i < len(items); i++ {
		f(i, items[i])
		//acumulador[i] = items[i] * 2
	}

	//return acumulador
}

func forEach(items []int, f func(int, int)) {
	for i := 0; i < len(items); i++ {
		f(i, items[i])
	}
}

func reduce(items []int, f func(int, int) int) int {
	acumulador := 0

	for i := 0; i < len(items); i++ {
		acumulador = f(acumulador, items[i])
	}

	return acumulador
}

func sumar(items []int) int {
	acumulador := 0

	for i := 0; i < len(items); i++ {
		acumulador = acumulador + items[i]
	}

	return acumulador
}

func duplicar(items []int) []int {
	acumulador := make([]int, len(items))

	for i := 0; i < len(items); i++ {
		acumulador[i] = items[i] * 2
	}
	return acumulador
}

func saludos() {
	saludar("nico")
	saludar("Gise")
	saludar("Saigo")

}

func saludar(nombre string) {
	fmt.Println("Hola como estas " + nombre + "!")
}
