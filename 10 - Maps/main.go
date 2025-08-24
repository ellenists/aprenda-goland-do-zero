package main

import (
	"fmt"
)

func main() {
	// MAPS
	m := map[string]int{ // tipo das chaves, tipo dos valores
		"idade": 30,
	}
	fmt.Println(m)
	fmt.Println(m["idade"])

	// Map aninhado dentro de outro Map
	n := map[int]map[string]string{
		1: {
			"Jane": "Doe",
		},
	}
	fmt.Println(n)
	fmt.Println(n[1])
	fmt.Println(n[1]["Jane"])

	// Remover uma determinada chave de um Map
	delete(m, "idade")
	fmt.Println(m)

	// Adicionar um novo item em um Map
	m["altura"] = 170
	fmt.Println(m)
}
