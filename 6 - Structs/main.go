package main

import "fmt"

// Um struct é um tipo de dado composto que agrupa zero ou mais valores de diferentes tipos em um único tipo.
// É usado para organizar dados relacionados, permitindo que você defina uma coleção de campos com nomes e tipos específicos.
// São úteis para representar entidades ou objetos.
type pokemon struct {
	id       uint8
	name     string
	type_    string
	defenses defenses
}

type defenses struct {
	normal   float32
	fire     float32
	water    float32
	electric float32
	grass    float32
	ice      float32
	fighting float32
	poison   float32
	ground   float32
	flying   float32
	psychic  float32
	bug      float32
	rock     float32
	ghost    float32
	dragon   float32
	dark     float32
	steel    float32
	fairy    float32
}

func main() {
	// 1: declarando um struct usando inferência de tipo (menos verboso).
	squirtle := pokemon{
		7,
		"Squirtle",
		"water",
		defenses{
			fire:     0.5,
			water:    0.5,
			electric: 2,
			grass:    2,
			ice:      0.5,
			steel:    0.5,
		},
	}
	fmt.Printf("%+v\n", squirtle) // %+v na função Printf exibe os nomes dos campos junto com seus valores.

	// 2: declarando um struct sem usar inferência de tipo (mais verboso).
	var wartortle pokemon
	wartortle.id = 8
	wartortle.name = "Wartortle"
	wartortle.type_ = "water"
	var defenses defenses
	defenses.electric = 2
	wartortle.defenses = defenses
	fmt.Printf("%+v\n", wartortle)
	wartortle.defenses.grass = 2
	fmt.Printf("%+v\n", wartortle)

	// 3: declarando um struct sem preencher todos os seus dados.
	// Para fazer isso, precisamos informar o nome dos campos que estamos informando.
	// Os campos que não forem informados terão o valor zero para o seu tipo.
	blastoise := pokemon{id: 9, name: "Blastoise", type_: "water"}
	fmt.Printf("%+v\n", blastoise)
}
