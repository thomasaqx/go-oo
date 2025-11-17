package main

import (
	"fmt"
	// importando package contabancaria

	"go-oo/clientes"
	"go-oo/conta"
)

func PagarBoleto(conta conta.VerificarConta, valorDoBoleto float64) {
    conta.Sacar(valorDoBoleto)
}



func main() {

	clienteThomas := clientes.Titular{
		Nome:      "Thomas",
		CPF:       "111.111.111-11",
		Profissao: "Desenvolvedor",
	}

	contaPThomas := conta.ContaPoupanca{
		Titular: clienteThomas,
		NumeroAgencia: 454,
		NumeroConta: 434453,
		Operacao: 1,
	}

	contaPThomas.Depositar(20000)
	fmt.Println("o saldo atual é: ", contaPThomas.ObterSaldo())

	PagarBoleto(&contaPThomas, 60)

	contaCThomas := conta.ContaCorrente{
		Titular: clienteThomas,
		NumeroAgencia: 454,
		NumeroConta: 434453,
	}

	contaCThomas.Depositar(12000)
	fmt.Println("o saldo atual da conta corrente é: ", contaCThomas.ObterSaldo())
	
	PagarBoleto(&contaCThomas, 500)










}
