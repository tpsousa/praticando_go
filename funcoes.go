


package main

import (
	"fmt"
)

func SemParametro() String{
	return "Exemplo de funcao sem parametro"
}

func umParametro(texto String) string{
	return texto
}

func doisParametros(texto string,numero int) (string,int){
	return texto , numero
}

func main () {
	fmt.Println(SemParametro())
	fmt.printLn(umParametro("exemplo de funcao com um parametro"))

	fmt.Println(doisParametros("passando dois parametros: essa string e o numero", 10))
}
    
// funcoes variadicas: 

 func somando(numeros ...int ) int {
	resultadoDaSoma := 0
	for _, numero := range numeros{
		resultadoDaSoma += numero
	}
	return resultadoDaSoma
 
}

func main (){
	fmt.Println(somando(1))
	fmt.Println(somando(1,1))
	fmt.Println(somando(1,1,1))
	fmt.Println(somando(1,1,2,4))
} 