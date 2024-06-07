package main

import (
	"fmt"
	"golang_poo/domains/models/contas"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDenis := contas.ContaPoupanca{}
	contaDenis.Depositar(100)
	PagarBoleto(&contaDenis, 60)

	fmt.Println(contaDenis.ObterSaldo())

}
