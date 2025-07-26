package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main.")
	auxiliar.Escrever()
	// Para referenciar um pacote usamos o nome presente após a última barra
	err := checkmail.ValidateFormat("test@mail.com")
	fmt.Println(err)
}
