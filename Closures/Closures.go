package Closures

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

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func Closures() {
	underRed.Print("\n\n\n Exemplo 14  (Closures) \n\n")
	underBlue.Printf(Explain.ExplainThis(14))
	fmt.Print("\n Aqui, usamos range para somar os números em um Slice. Arrays funcionam assim também. \n")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	package main

	import "fmt"
	`)
	underBlue.Print("\n A função intSeq abaixo retorna outra função, que definimos anonimamente no corpo da função intSeq. A função retornada fecha sobre a variável ipara formar um fechamento.\n")

	exampleStyle.Print(`
	func intSeq() func() int {
		i := 0
		return func() int {
			i++
			return i
		}
	}

	func main() {
	`)
	underBlue.Print("\n Chamamos a função intSeq, atribuindo o resultado (uma função) a nextInt. Este valor de função captura seu próprio valor de i, que será atualizado cada vez que chamarmos nextInt.\n")
	exampleStyle.Print(`
		nextInt := intSeq()
	`)
	underBlue.Print("\n Veja o efeito do fechamento chamando nextInt algumas vezes.\n")
	exampleStyle.Print(`
		fmt.Println(nextInt())
		fmt.Println(nextInt())
		fmt.Println(nextInt())
	`)
	underBlue.Print("\n Para confirmar que o estado é exclusivo para aquela função específica, crie e teste um novo.\n")
	exampleStyle.Print(`
		newInts := intSeq()
		fmt.Println(newInts())

	}
	`)
	resultText.Print("\nRESULTADO : \n")
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInts := intSeq()
	fmt.Println(newInts())

}
