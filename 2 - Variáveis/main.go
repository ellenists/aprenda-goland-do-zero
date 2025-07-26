package main

import (
	"fmt"
)

func main() {
	// Existem algumas maneiras de declarar variáveis e constantes em go.
	// A primeira delas é explicitando o tipo da variável:
	const a string = "a"
	// Há também a possibilidade de deixar que o go infira o tipo da variável:
	b := "b"

	// Para declarar mais de uma variável de uma vez, podemos usar:
	var (
		c string
		d int
	)

	c = "c"
	d = 1

	// O mesmo vale para constantes:
	const (
		e string = "e"
		f int    = 2
	)

	g, h := "g", 3
	fmt.Println(a, b, c, d, e, f, g, h)

	// Um ponto curioso de go, é que não precisamos criar uma variável auxiliar para
	// trocar o valor entre duas variáveis, podemos fazer diretamente:
	fmt.Println(d, h)
	d, h = h, d
	fmt.Println(d, h)
}
