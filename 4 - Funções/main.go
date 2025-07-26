package main

import (
	"fmt"
	"strings"
)

// se dois ou mais parâmetros são do mesmo tipo, conseguimos declarar este tipo somente uma vez
func soma(n1, n2 int8) int8 {
	return n1 + n2
}

// Uma função pode ter mais de um valor retornado:
func caseUnifier(text string) (string, string) {
	lowercase := strings.ToLower(text)
	uppercase := strings.ToUpper(text)
	return lowercase, uppercase
}

func main() {
	resultado := soma(2, 2)
	fmt.Println(resultado)

	// Funções podem ser atribuídas a variáveis:
	var a = func(text string) string {
		return text
	}
	retorno := a("blah")
	fmt.Println(retorno)

	// Se quiser ignorar algum dos valores retornados pela função, use o blank identifier _:
	// lower, _ := caseUnifier("pApAGAiO")
	lower, upper := caseUnifier("pApAGAiO")
	fmt.Println(lower, upper)

}
