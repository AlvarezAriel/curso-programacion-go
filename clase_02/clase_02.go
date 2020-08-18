package main

import "fmt"

func estaLaPalabraEnLaFrase(palabra string, frase string) bool {
	var index int = 0
	var size int = len(frase)

	for index < size {

		if palabraEstaEnPosicion(palabra, frase, index) {
			return true
		}

		index = index + 1
	}

	return false
}

func palabraEstaEnPosicion(palabra string, frase string, offset int) bool {
	largoPalabra := len(palabra)
	largoFrase := len(frase) - offset

	if largoPalabra > largoFrase  {
		return false
	}

	var index int = 0

	for index < largoPalabra {
		if palabra[index] != frase[index + offset] {
			return false
		}
		index = index + 1
	}

	return true
}

func main() {

	palabraQueBusco := "yo no los conozco"
	frase :=           "son los orozco, yo los conozco, son ocho los monos"

	fmt.Println(estaLaPalabraEnLaFrase(palabraQueBusco, frase))
}


