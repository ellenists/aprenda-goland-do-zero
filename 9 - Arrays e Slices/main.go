package main

import (
	"fmt"
	"reflect"
)

func main() {
	// ARRAYS
	// possuem um tamanho fixo, por isso são menos flexíveis
	// aloca memória fixa para todos os elementos no momento da criação
	// é passado para funções por cópia, o que pode ser ineficiente para arrays grandes, pois toda a estrutura é duplicada
	// não pode ser redimensionado.
	fmt.Println("ARRAYS")

	// primeira forma de declarar
	// usando a cláusula var
	var a [5]int
	a[0] = 1
	fmt.Println(a)

	// segunda forma de declarar
	// usando inferência de tipo
	// e já podemos inserir valores no array diretamente
	b := [5]int{1, 2, 3, 4}
	fmt.Println(b)
	b[4] = 5
	fmt.Println(b)

	// terceira forma de declarar
	// usando inferência de tipo
	// e deixando que o go determine o tamanho do array
	// de acordo com a quantidade de itens que armazenamos nele
	// nesse caso, o array terá tamanho 5 porque estou armazenando 5 valores nele
	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println(c)
	fmt.Println(len(c))

	fmt.Println(reflect.TypeOf(c))

	// Como o tamanho de um array é fixo, não há uma operação nativa para remover elementos de um array em Go
	// No entanto, podemos substituir um elemento atribuindo um valor (como zero) ao índice desejado
	// ou copiar os elementos desejados para um novo array, excluindo o que você quer remover

	// SLICES
	// Mais flexível, pois pode crescer com append (até a capacidade)
	// tem um tamanho dinâmico que é determinado pelo número de elementos atualmente armazenados
	// É uma view de um array subjacente, armazenando um ponteiro para o array, a memória é gerenciada dinamicamente
	// é passado para funções por referência (o ponteiro é copiado), tornando-o mais eficiente para manipulações

	fmt.Println("SLICES")

	d := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(d)
	fmt.Println(reflect.TypeOf(d))

	d = append(d, 10)
	fmt.Println(d)

	e := c[0:3] // 0 inclusive e 3 exclusive
	fmt.Println(e)

	c[0] = 10000
	fmt.Println(e)

	// Em Go não existe uma função built-in específica como remove ou delete para slices
	// Para "remover" um elemento de um slice, você basicamente recria o slice sem o elemento desejado
}
