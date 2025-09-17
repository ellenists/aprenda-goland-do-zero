package main

import "fmt"

func divisao(n1, n2 int) (resultado, resto int) {
	resultado = n1 / n2
	resto = n1 % n2
	return
}

func main() {
	fmt.Print(divisao(5, 2))
}
