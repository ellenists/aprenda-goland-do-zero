package main

import "fmt"

func main() {
	a := 10
	if a > 10 {
		fmt.Println("Maior que 10.")
	} else if a == 10 {
		fmt.Println("Igual a 10.")
	} else {
		fmt.Println("Menor que 10.")
	}

	// Em Go, o "if init" é uma feature que permite declarar e inicializar
	// uma variável local dentro da própria cláusula if.
	// Essa variável só existe dentro do escopo do if (e do else, se houver)
	// e pode ser usada para preparar dados antes de avaliar a condição.
	if b := 10; b < 20 {
		fmt.Printf("%d é menor que 20.", b)
	} else if b == 20 {
		fmt.Printf("%d é igual a 20.", b)
	} else {
		fmt.Printf("%d é maior que 20.", b)
	}

	// fmt.Println(b) não funciona porque a variável b só existe dentro do escopo if/else
}
