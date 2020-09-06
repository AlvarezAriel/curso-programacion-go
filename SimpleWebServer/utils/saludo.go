package utils

import("strings")

func Saludar() string {
	return ponerleEnfasis("Hola", 3)
}

func ponerleEnfasis(frase string, nivelDeEnfasis int) string {
	return frase + strings.Repeat("!", nivelDeEnfasis)
}
