package main

import (
	"fmt"
)

func main() {
	var numero int = 42
	ponteiroNumero := &numero

	// Tentativa de atribuir um valor inteiro diretamente ao ponteiro
	// Isso causará um erro de compilação
	ponteiroNumero = 200 // Esta linha está comentada porque causaria um erro

	fmt.Println("Valor original:", *ponteiroNumero) // 42
}
