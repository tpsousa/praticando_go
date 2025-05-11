package main

import (
  "fmt"
  "github.com/tpsousa/praticando_go/contas"
)


func main() {
	// Criamos uma conta com saldo inicial.
	conta := contas.ContaCorrente{titular: "Thiago", saldo: 10}

	// Realizamos um depósito e mostramos o saldo.
	status := conta.Depositar(50)
	fmt.Println(status)
	fmt.Println("Saldo da conta do Thiago:", conta.saldo)

	// Criamos outra conta.
	contaDaSilvia := ContaCorrente{titular: "Silvia", saldo: 500}
	fmt.Println("Saldo da Silvia antes do depósito:", contaDaSilvia.saldo)

	// Fazemos um depósito na conta da Silvia.
	status = contaDaSilvia.Depositar(2000)
	fmt.Println(status)
	fmt.Println("Saldo da Silvia após o depósito:", contaDaSilvia.saldo)

	// Criamos mais uma conta.
	contaDoGustavo := ContaCorrente{titular: "Gustavo", saldo: 100}
	fmt.Println("Saldo do Gustavo antes da transferência:", contaDoGustavo.saldo)

	// Fazemos uma transferência da Silvia para o Gustavo.
	statusTransferencia := contaDaSilvia.Transferir(200, &contaDoGustavo)

	// Mostramos o resultado da transferência e os saldos atualizados.
	fmt.Println("Transferência realizada?", statusTransferencia)
	fmt.Println("Saldo da Silvia após transferência:", contaDaSilvia.saldo)
	fmt.Println("Saldo do Gustavo após receber:", contaDoGustavo.saldo)
}
