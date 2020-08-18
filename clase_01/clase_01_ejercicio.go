package main

import "fmt"

func main() {
	cuantasLetrasTiene("Gisele")
	cuantasLetrasTiene("Ema")
	cuantasLetrasTiene("Tsuki")
	cuantasLetrasTiene("Gianni")
	cuantasLetrasTiene("Eric")
}

func cuantasLetrasTiene(nombre string) {
	var saludo string = "Hola " + nombre

	if esStringPar(nombre) {
		saludo = saludo + "!"
	}

	fmt.Println(saludo)
}

func esStringPar(texto string) bool {
	return esNumeroPar(len(texto))
}

func esNumeroPar(unNumero int) bool {
	return 0 == unNumero % 2
}
