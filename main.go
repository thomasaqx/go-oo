package main

import (
	"fmt"
	// importando package contabancaria

	"go-oo/clientes"
	"go-oo/conta"
)

func main() {

	clienteThomas := clientes.Titular{
		Nome:      "Thomas",
		CPF:       "111.111.111-11",
		Profissao: "Desenvolvedor",
	}

	contaThomas := conta.ContaCorrente{
		Titular: clienteThomas,
		NumeroAgencia: 454,
		NumeroConta: 434453,
	}

	contaThomas.Depositar(2000)
	fmt.Println("o saldo atual é: ", contaThomas.ObterSaldo())




	fmt.Println(contaThomas)
}
