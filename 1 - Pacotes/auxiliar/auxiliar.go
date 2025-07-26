package auxiliar

import (
	"fmt"
)

// Escrever exibe determinada mensagem no console.
// É boa prática adicionar comentários descrevendo funções exportadas.
// Se uma função começa com letra maiúscula, ela está sendo exportada, ou seja, estará visível para outros pacotes da aplicação.
// Da mesma forma, se começa com letra minúscula ela fica visível somente dentro do pacote onde se encontra.
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar.")
}
