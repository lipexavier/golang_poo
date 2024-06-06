package contas

import "golang_poo/domains/models/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente!"
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.Saldo += valorDeposito
		return "Depósito realizado com sucesso!", c.Saldo
	} else {
		return "Valor do depósito é inválido!", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia > 0 && valorTransferencia <= c.Saldo {
		c.Saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}
