package main

import "fmt"

type ContaCorrente struct {
	titular string
	saldo   float64
}

// Este é um método chamado Depositar.
// Ele é associado à struct ContaCorrente por meio do receiver (c *ContaCorrente),
// ou seja, é como se fosse um "método de instância".
// O receiver é um ponteiro, permitindo modificar o saldo original.
// O método recebe um valor e adiciona ao saldo, retornando uma mensagem de status.
func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso"
	} else {
		return "Valor do depósito menor que zero"
	}
}

// Método para transferir valor de uma conta para outra.
// Retorna verdadeiro se a transferência foi bem-sucedida.
func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia > 0 && valorDaTransferencia <= c.saldo {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func main() {
	// Criamos uma conta com saldo inicial.
	conta := ContaCorrente{titular: "Thiago", saldo: 10}

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
