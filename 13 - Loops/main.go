package main

import (
	"fmt"
	"strings"
)

func main() {
	// primeiro uso, o mais simples:
	i := 0
	for i < 10 {
		i++
		fmt.Println(i)
	}

	// ainda primeiro uso, mas um loop infinito:
	// for {
	// 	fmt.Println("executando infinitamente...")
	// 	time.Sleep(time.Second)
	// }

	// segundo uso, mais convencional:
	for j := 0; j < 10; j += 2 {
		// a variável j existirá somente no escopo do loop
		fmt.Println(j)
	}

	// terceiro uso, loop sobre um iterável (range):
	pokemons := [3]string{"Bulbasaur", "Charmander", "Squirtle"}
	// se não precisar usar o indice, substituí-lo por _ (blank identifier)
	for indice, pokemon := range pokemons {
		fmt.Println(indice, pokemon)
	}

	// ainda no terceiro uso, exemplo de iteração sobre uma string:
	for _, letra := range "blastoise" {
		// aqui convertemos a letra em string porque, por padrão
		// go referencia o valor correspondente a determinada letra
		// na tabela ASCII; então para obter a letra em si, convertemos.
		fmt.Print(strings.ToUpper(string(letra)))
	}

	// exemplo do terceiro uso iterando sobre um map:
	pokemons2 := map[uint8]string{
		1: "Bulbasaur",
		2: "Charmander",
		3: "Squirtle",
	}

	// Em Go, mapas (map) não garantem ordem de iteração.
	// Ou seja: quando você faz um for range em pokemons2,
	// a ordem em que as chaves/valores aparecem pode variar
	// a cada execução do programa.
	for i, pokemon := range pokemons2 {
		fmt.Printf("\nPokémon de número %d: %s", i, pokemon)
	}

}
