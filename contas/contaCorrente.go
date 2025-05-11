

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
