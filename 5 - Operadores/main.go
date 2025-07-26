package main

import "fmt"

func main() {
	// OPERADORES ARITMÉTICOS
	soma := 2 + 2
	subtracao := 2 - 2
	multiplicacao := 2 * 2
	divisao := 2 / 2
	resto := 2 % 2
	fmt.Println(soma, subtracao, multiplicacao, divisao, resto)

	// Go não permite operações entre variáveis de tipos distintos!

	// OPERADORES DE ATRIBUIÇÃO
	var a int = 1
	b := 2
	fmt.Println(a, b)

	// OPERADORES RELACIONAIS
	// Sempre retornam um valor boleano.
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)
	fmt.Println(1 == 2)

	// OPERADORES LÓGICOS
	fmt.Println(true && false) // e
	fmt.Println(true || false) // ou
	fmt.Println(!true)         // negação

	// OPERADORES UNÁRIOS
	num := 10
	num++ // num = num + 1
	fmt.Println(num)
	num += 4 // num = num + 4
	fmt.Println(num)
	num--    // num = num - 1
	num *= 2 // num = num * 2
	num /= 2 // num = num / 2
	num %= 3 // num = num % 2
	fmt.Println(num)

	// O OPERADOR TERNÁRIO NÃO ESTÁ DISPONÍVEL EM GO
	// podemos usar o bom e velho if else para cumprir sua função.
	// c := num > 3 ? "Maior que 3" : "Menor que 3"
	var c string
	if num > 3 {
		c = "Maior que 3"
	} else {
		c = "Menor que 3"
	}
	fmt.Println(c)

}
