package main

import "fmt"

type Golondrina struct {
	nombre  string
	energia float64
	peso    float64
}

func (golondrina *Golondrina) volar(km float64) {
	golondrina.energia -= km * 1.3
}

func comer(golondrina *Golondrina, kgComida float64) {
	golondrina.energia += kgComida * 0.8
}

type Paloma struct {
	nombre  string
	energia float64
	peso    float64
}

func (paloma *Paloma) volar(km float64) {
	paloma.energia -= km * 1.3 * (paloma.peso * 0.1)
}

func robar(paloma *Paloma, kgComida float64) {
	paloma.energia += kgComida * 0.1
}

type Volador interface {
	volar(km float64)
}

func main() {
	pepita := Golondrina{nombre: "Pepita"}

	comer(&pepita, 5)
	pepita.volar(2)

	fmt.Println(pepita)

	palomon := Paloma{nombre: "Palomon", peso: 4}

	robar(&palomon, 100)

	palomon.volar(2)

	fmt.Println(palomon)

	volarMucho(&palomon, &pepita)

	fmt.Println(pepita)
	fmt.Println(palomon)
}

func volarMucho(voladores ...Volador) {
	for i := 0; i < len(voladores); i++ {
		voladores[i].volar(100)
	}
}

func ejemplo1() {

	nico := Inscripto{
		nombre:               "Nico",
		edad:                 18,
		esPrimeraInscripcion: true,
	}

	ariel := Inscripto{
		nombre:               "Ariel",
		edad:                 19,
		esPrimeraInscripcion: false,
	}

	inscriptos := []Inscripto{nico, ariel}

	editarNombre(&inscriptos[1], "Logain")

	listarInscriptos(inscriptos)
}

type Inscripto struct {
	nombre               string
	edad                 int
	esPrimeraInscripcion bool
}

func editarNombre(inscripto *Inscripto, nuevoNombre string) {
	inscripto.nombre = nuevoNombre
	fmt.Println("adentro de editarNombre: ", inscripto.nombre, inscripto.edad)
}

func listarInscriptos(inscriptos []Inscripto) {
	for i := 0; i < len(inscriptos); i++ {
		inscripto := inscriptos[i]
		fmt.Println(inscripto.nombre, inscripto.edad)
	}
}
