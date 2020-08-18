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

	var largoDelNombre int = len(nombre)
	var esPar bool = 0 == largoDelNombre%2

	if esPar {
		saludo = saludo + "!"
	}

	fmt.Println(saludo)
}
