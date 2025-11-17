package main

import (
	"fmt"
)

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	contaThomas := ContaCorrente{
		"Thomas A. Queiroz",
		123456,
		5436,
		2150.50,
	}

	fmt.Println(contaThomas)

}
