package main

import "fmt"

type ContaCorrente struct {
	titular string
	saldo   float64
}

// Este é um método chamado Depositar.
// Ele é associado à struct ContaCorrente por meio do receiver (c *ContaCorrente),
// ou seja, é como se fosse um "método de instância".
// O parâmetro c é um ponteiro para uma instância de ContaCorrente,
// permitindo modificar o saldo da conta original.
// O método recebe um valor de depósito e o adiciona ao saldo da conta.
func (c *ContaCorrente) Depositar(valorDoDeposito float64) {
    if valorDoDeposito > 0 {
        c.saldo += valordoDeposito
        return "Deposito realizado com sucesso"
    } else { 
        return "Valor do depósito menor que zero"
    }
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64,contaDestino *ContaCorrente) bool{
	if valorDaTransferencia < c.saldo && valorDaTransferencia> 0{
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	}else {
		return false
	}
}

func main() {
	// Criamos uma variável chamada conta que é uma instância da struct ContaCorrente,
	// com os campos titular e saldo inicializados.
	conta := ContaCorrente{titular: "Thiago", saldo: 10}

	// Chamamos o método Depositar na variável conta para adicionar 50 ao saldo.
	conta.Depositar(50)

	// Imprimimos o novo saldo da conta.
	fmt.Println(conta.saldo) // Correção aqui: era "printLn", o correto é "Println" com "P" maiúsculo.

	contaDaSilvia := contaContaCorrente{}
    contaDaSilvia.titular = "Silvia"
    contaDaSilvia.saldo = 500
    
	fmt.Println(contaDaSilvia.saldo) 
    status, valor := contaDaSilvia.Depositar(2000)
    fmt.Println(status, valor) 

	contaDoGustavo := ContaCorrente{titular: "gustavo", saldo: 100}

	fmt.printLn(contaDaSilvia)
	fmt.Println(contaDoGustavo)

	status := contaDaSilvia.Transferir(200,&contaDoGustavo)

	fmt.print(Status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

}
