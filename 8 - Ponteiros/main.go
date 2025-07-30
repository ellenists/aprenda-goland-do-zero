package main

import "fmt"

func main() {
	// O fato de a variável a ter sido incrementada não afeta em nada
	// o conteúdo da variável b, pois b recebeu uma CÓPIA do valor de a.
	a := 1
	b := a
	fmt.Println(a, b)
	a++
	fmt.Println(a, b)

	// O fato de a variável c ter sido incrementada afeta também o conteúdo da variável d,
	// porque d está apontando para o ENDEREÇO DE MEMÓRIA (&c) onde o conteúdo de c está armazenado.
	// Sendo assim, se alteramos o que está armazenado neste endereço, todo ponteiro que estiver
	// apontando para ele refletirá a alteração.
	c := 10
	d := &c
	fmt.Println(c, d)
	c++
	fmt.Println(c, d)
	// O asterisco antes da variável (*d) promove o dereferencing
	// que é o ato de acessar o valor armazenado no endereço de memória apontado por um ponteiro.
	// Então, na instrução abaixo, o console exibe o valor armazenado na variável ao invés do endereço de memória.
	fmt.Println(c, *d)
}
