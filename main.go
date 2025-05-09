


import fmt

//se precisarmos usar apenas alguns,teremos que especificar o nome da classe/ variavel adequada
func main () {
	contaDoGuilherme := ContaCorrente{titular : "Guilherme ",numeroAgencia: 589, numeroConta: 123456,saldo: 125.5 }
	fmt.printLn(contaDoGuilherme)


	contaDaBruna := ContaCorrente{"Bruna", 222,111222,200}

	fmt.printLn(contaDoGuilherme)
	fmt.printLn(contaDaBruna)

	var contaDaCris *ContaCorrente

	contaDaCris = new(ContaCorrente)
    //acessando atributos de uma classe e atribuindo valores para ela
	contaDaCris.titular = "Cris"

	fmt.printLn(*contaDaCris)
}