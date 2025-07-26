package main

import (
	"errors"
	"fmt"
)

func main() {
	// NÚMEROS INTEIROS
	// Signed int: int (suportam números negativos).
	// Unsigned int: uint (NÃO suportam números negativos).
	// Variam de acordo com o tamanho do número que armazenarão:
	// int8: 8 bits (1 byte), intervalo: -128 a 127.
	// int16: 16 bits (2 bytes), intervalo: -32.768 a 32.767.
	// int32: 32 bits (4 bytes), intervalo: -2.147.483.648 a 2.147.483.647.
	// int64: 64 bits (8 bytes), intervalo: -9.223.372.036.854.775.808 a 9.223.372.036.854.775.807.
	// int: tamanho depende da arquitetura (32 bits em sistemas 32-bit, 64 bits em sistemas 64-bit).
	var foo int = -1
	var bar uint = 1
	fmt.Println(foo, bar)

	// Aliases
	var a int32
	// é o mesmo que:
	var b rune

	// assim como:
	var c uint8
	// é equivalente a:
	var d byte

	fmt.Println(a, b, c, d)

	// NÚMEROS REAIS
	// float32 ou float64, como nos inteiros, o que difere é o tamanho que suportam
	// Se usarmos inferência de tipos, a arquitetura do sistema será usada para definir o tamanho.
	var e float32 = 123.45
	var f float64 = 12345.67
	g := 123.45

	fmt.Println(e, f, g)

	// STRINGS
	// cadeia de caracteres, usamos sempre aspas duplas
	var h string = "h"
	i := "i"
	fmt.Println(h, i)

	// Go não suporta char como estamos acostumados, ao invés disso
	// ele os interpreta usando o valor equivalente dele na tabela ASCII.
	char := 'J'
	fmt.Println(char) // printa 74 no console

	// BOOLEAN
	var booleano bool = true
	booleano2 := false
	fmt.Println(booleano, booleano2)

	// ERROR
	// é um tipo próprio do go, seu valor zero é nil.
	var erro error = errors.New("error message")
	fmt.Println(erro)

	// VALORES ZERO
	// é o valor padrão atribuído a uma variável quando a mesma não é inicializada.
	// https://go.dev/tour/basics/12
	var k int
	var l float64
	var n bool
	var s string
	var p *int
	var sl []int
	var m map[string]int
	var o chan int
	var fn func()
	var iface interface{}

	fmt.Println("Zero Values:")
	fmt.Println("int:", k)           // 0
	fmt.Println("float64:", l)       // 0
	fmt.Println("bool:", n)          // false
	fmt.Println("string:", s)        // ""
	fmt.Println("pointer:", p)       // <nil>
	fmt.Println("slice:", sl)        // []
	fmt.Println("map:", m)           // map[]
	fmt.Println("channel:", o)       // <nil>
	fmt.Println("function:", fn)     // <nil>
	fmt.Println("interface:", iface) // <nil>
}
