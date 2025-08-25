package main

import "fmt"

// Essa é a maneira mais comum e mais simples
func getColor1(code string) string {
	switch code {
	case "R":
		return "RED"
	case "G":
		return "GREEN"
	case "B":
		return "BLUE"
	default:
		return "Invalid Code!"
	}
}

// Essa é uma outra forma, mais flexível
func getColor2(code string) string {
	switch {
	case code == "R":
		return "RED"
	case code == "G":
		return "GREEN"
	case code == "B":
		return "BLUE"
	default:
		return "Invalid Code!"
	}
}

// Aqui usamos a cláusula fallthrough em uma estrutura switch
// Ela é usada para forçar a execução do caso seguinte, mesmo que a condição do caso atual
// já tenha sido satisfeita.
// Por padrão, após um caso ser executado, o controle sai do switch e a cláusula fallthrough
// impede isso; é pouco usada mas interessante saber que existe.
func getColor3(code string) string {
	var color string
	switch code {
	case "R":
		color = "RED"
	case "G":
		color = "GREEN"
		fallthrough
	case "B":
		color = "BLUE"
	default:
		color = "Invalid Code!"
	}
	return color
}

func main() {
	// SWITCH
	// Existem duas formas de usar essa cláusula e não precisamos usar 'break'
	fmt.Println(getColor1("R"))
	fmt.Println(getColor2("G"))
	fmt.Println(getColor3("G"))
}
