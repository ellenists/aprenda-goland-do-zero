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

	// ARRAYS INTERNOS
	// Em Go, um slice não é um array por si só, mas sim uma view de um array subjacente.
	// Quando você cria um slice, ele é baseado em um array interno que pode ser compartilhado por outros slices.
	// O array subjacente é alocado na memória, e o slice apenas "aponta" para uma porção desse array.

	// Compartilhamento: Vários slices podem apontar para o mesmo array subjacente.
	// Alterações em um slice afetam o array original e, indiretamente, outros slices que o compartilham.

	// Capacidade: A capacidade (cap) é o número de elementos disponíveis a partir do início do slice até o final do array subjacente.
	// Isso determina até onde o slice pode crescer com append antes de precisar realocar.

	// Realocação: Se você usa append e excede a capacidade, Go cria um novo array subjacente maior e copia os dados, atualizando o ponteiro do slice.
	fmt.Println("-------------------------------")
	f := make([]float32, 10) // tipo, tamanho, capacidade máxima
	// A capacidade máxima é opcional, se omitida ela assumirá o valor do parâmetro tamanho
	// Então, podemos criar o slice também desta forma: make([]float32, 10, 12)
	fmt.Println(f)
	fmt.Println(len(f))
	fmt.Println(cap(f))

	// Ao adicionar mais 3 itens no slice que foi definido como tendo capacidade para 12 posições
	// ele ultrapassa essa capacidade; o Go lida com isso criando um novo array de tamanho dobrado
	// e apontando o slice para ele.
	f = append(f, 0)    // 11 itens no slice, o que ultrapassa sua capacidade de 10
	fmt.Println(len(f)) // seu tamanho passa a ser 11
	fmt.Println(cap(f)) // sua capacidade passa a ser 20: a capacidade anterior (10) multiplicada por 2
}
