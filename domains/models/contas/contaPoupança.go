package contas

import "golang_poo/domains/models/clientes"

type ContaPoupanca struct {
	Titular                          clientes.Titular
	Numero, Agencia, Conta, Operacao int
	saldo                            float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "saldo insuficiente!"
	}
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Depósito realizado com sucesso!", c.saldo
	} else {
		return "Valor do depósito é inválido!", c.saldo
	}
}

func (c *ContaPoupanca) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia > 0 && valorTransferencia <= c.saldo {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo

}
