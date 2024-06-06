package models

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
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insufiiente!"
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.Saldo += valorDeposito
		return "Deposito realizado com sucesso!", c.Saldo
	} else {
		return "Valor do deposito é inválido!", c.Saldo
	}

}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaTransferencia *ContaCorrente) bool {
	if valorTransferencia < c.Saldo && valorTransferencia > 0 {
		contaTransferencia.Depositar(valorTransferencia)
		c.Sacar(valorTransferencia)
		return true
	} else {
		return false
	}
}
