package Pointers

import (
	"fmt"

	Explain "github.com/AndreoBouzas/Go_By_Example/Explain"
	color "github.com/fatih/color"
)

var (
	underBlue    = color.New(color.FgCyan)
	underRed     = color.New(color.FgHiRed).Add(color.Underline)
	exampleStyle = color.New(color.FgHiWhite)
	exampleText  = color.New(color.FgHiGreen)
	resultText   = color.New(color.FgHiMagenta)
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func Pointers() {
	underRed.Print("\n\n\n Exemplo 16  (Ponteiros) \n\n")
	underBlue.Printf(Explain.ExplainThis(16))
	fmt.Print("\n Mostraremos como os ponteiros funcionam em contraste com os valores com 2 funções: zeroval e zeroptr. \n ")
	exampleText.Print("\nEXEMPLO : \n")
	fmt.Print("\n zeroval tem um parâmetro int, então os argumentos serão passados ​​a ele por valor.\n")
	fmt.Print(" zeroval obterá uma cópia de ival distinto daquele na função de chamada.\n")
	exampleStyle.Print(`
	package main
	
	import "fmt"
	
	func zeroval(ival int) {
		ival = 0
	}
	`)
	underBlue.Print("\n zeroptr em contraste, tem um parâmetro *int, o que significa que leva um ponteiro int.\n")
	fmt.Print(" O código *iptr no corpo da função cancela a referência do ponteiro de seu endereço de memória para o valor atual naquele endereço.\n")
	fmt.Print(" Atribuir um valor a um ponteiro não referenciado altera o valor no endereço referenciado.\n")

	exampleStyle.Print(`
	func zeroptr(iptr *int) {
		*iptr = 0
	}
	
	func main() {
		i := 1
		fmt.Println("initial:", i)
		zeroval(i)
		fmt.Println("zeroval:", i)
	`)
	underBlue.Print("\n A sintaxe &i fornece o endereço de memória de i, ou seja, um ponteiro para i.\n")

	exampleStyle.Print(`
		zeroptr(&i)
		fmt.Println("zeroptr:", i)
	`)
	underBlue.Print("\n Ponteiros também podem ser impressos.\n")

	exampleStyle.Print(`
		fmt.Println("pointer:", &i)
	}
	`)
	underBlue.Print("\n zeroval não artera o i na função main, más zeroptr altera, pois possui referencia ao endereço de memoria desta variável\n")
	resultText.Print("\nRESULTADO : \n")
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)

}
