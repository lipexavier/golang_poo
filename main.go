package main

import (
	"fmt"
	contas "golang_poo/domains/models"
)

func main() {

	contaDaSilvia := contas.ContaCorrente{}
	contaDaSilvia.Titular = "Silvia"
	contaDaSilvia.Saldo = 500

	contaDoGustavo := contas.ContaCorrente{}
	contaDoGustavo.Titular = "Gustavo"
	contaDoGustavo.Saldo = 1000

	fmt.Println(contaDaSilvia.Saldo)

	fmt.Println(contaDaSilvia.Sacar(100))
	fmt.Println(contaDaSilvia.Saldo)

	status, valor := contaDaSilvia.Depositar(2000)
	fmt.Println(status, valor)

	fmt.Println("Antes da transferência")
	fmt.Println("Saldo conta Gustavo: ", contaDoGustavo.Saldo)
	fmt.Println("Saldo conta Silvia: ", contaDaSilvia.Saldo)

	contaDoGustavo.Transferir(200, &contaDaSilvia)

	fmt.Println("Depois da transferência")
	fmt.Println("Saldo conta Gustavo: ", contaDoGustavo.Saldo)
	fmt.Println("Saldo conta Silvia: ", contaDaSilvia.Saldo)

}
