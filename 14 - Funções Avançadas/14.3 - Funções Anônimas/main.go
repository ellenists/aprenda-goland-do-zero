package main

import "fmt"

func main() {

	result := func(txt string) string {
		return fmt.Sprintf("Meu pokémon favorito é: %s", txt)
	}("Squirtle")

	fmt.Print(result)
}
