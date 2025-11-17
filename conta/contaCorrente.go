package conta

import "fmt"

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		fmt.Println("valor sacado: ", valorDoSaque)
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		fmt.Println("valor depositado: ", valorDoDeposito)
		return "Depósito realizado com sucesso"
	} else {
		return "Valor do depósito inválido"
	}
}
