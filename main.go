package main

import (
	"fmt"
	// importando package contabancaria

	"go-oo/conta"
)

func main() {
	// criando objeto usando construtor
	contaThomas := conta.ContaCorrente{
		Titular: "Thomas A. Queiroz",
		Saldo:   2150.50,
	}

fmt.Println(contaThomas.Depositar(255.0))

fmt.Println(contaThomas.Sacar(100.0))

fmt.Println("Status da conta: ", contaThomas )


}
