package main

import "fmt"

func soma(numeros ...int) int {
	t := 0
	for _, numero := range numeros {
		t += numero
	}
	return t
}

func escrever(s string, numeros ...int) {
	// uma função não pode ter mais de um parâmetro variático
	// mas pode ter vários parâmetros fixos
	// é esse parâmetro precisa obrigatoriamente ser o último a ser informado na função
	// escrever(numeros ...int, s string) não funcionaria.
	for _, numero := range numeros {
		fmt.Println(s, numero)
	}
}

func main() {
	// posso enviar quantos números quiser
	fmt.Println(soma(200, 4534, 1212, 19))
	// também posso não enviar parâmetro nenhum
	fmt.Println(soma())
	// e posso combinar parâmetros fixos com um variável
	escrever("foo", 5, 3, 8, 4, 1)
}
