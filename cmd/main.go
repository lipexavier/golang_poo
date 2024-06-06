package main

import (
	"fmt"
	"golang_poo/domains/models/clientes"
	"golang_poo/domains/models/contas"
)

func main() {
	clienteBruno := clientes.Titular{
		Nome:      "Bruno",
		CPF:       "123.111.123.12",
		Profissao: "Desenvolvedor"}

	contaDoBruno := contas.ContaCorrente{clienteBruno, 123, 123456, 100}

	fmt.Println(contaDoBruno)

}
