package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type IntExpression interface {
	exec() int
}

//--------- MULTIPLICACION -----------
type Mul struct {
	a, b IntExpression
}

func (op Mul) exec() int {
	return op.exec() * op.exec()
}

//--------- SUMA -----------
type Sum struct {
	a, b IntExpression
}

func (op Sum) exec() int {
	return op.exec() + op.exec()
}

//--------- LITERAL -----------
type Num struct {
	a int
}

//--------- DIVISION ENTERA -----------
type Div struct {
	a, b IntExpression
}

func (op Div) exec() int {
	return op.exec() / op.exec()
}

func (literal Num) exec() int {
	return literal.a
}

func main() {
	resultado, err := ejecutarCuenta("( 5 + 3 ) * 6")

	expresion := Mul{Sum{Num{5}, Num{3}}, Num{6}}

	fmt.Println("Resultado de expresion a mano: %d", expresion.exec())

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("El resultado es %d\n", resultado)
	}
}

func parse(input string) (IntExpression, error) {

	tokens := strings.Split(input, " ")

	a, err := strconv.Atoi(tokens[0])
	if err != nil {
		return nil, err
	}

	b, err := strconv.Atoi(tokens[2])
	if err != nil {
		return nil, err
	}

	expresion, err := opBuilder(tokens[1], a, b)
	if err != nil {
		return nil, err
	}

	return expresion, nil
}

func opBuilder(opString string, a int, b int) (IntExpression, error) {
	if opString == "*" {
		return Mul{a, b}, nil
	} else if opString == "+" {
		return Sum{a, b}, nil
	} else {
		return nil, errors.New(opString + " no es un operador valido")
	}
}

func ejecutarCuenta(cuentaStr string) (int, error) {

	operacion, err := parse(cuentaStr)
	if err != nil {
		return 0, err
	}

	return operacion.exec(), nil
}
