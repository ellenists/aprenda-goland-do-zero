package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint8
}

type estudante struct {
	// Uma struct (pessoa) é incluída como um campo aninhado em outra struct (estudante).
	// Isso se chama "embedding" (ou incorporação).
	pessoa
	id    uint8
	curso string
}

func main() {
	e1 := estudante{pessoa: pessoa{"Jane", 21}}
	fmt.Printf("%+v\n", e1)
	e1.idade = 22
	e1.id = 1
	e1.curso = "COMPSCI"
	fmt.Printf("%+v\n", e1)
}
