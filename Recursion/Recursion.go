package Recursion

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

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func Recursion() {
	underRed.Print("\n\n\n Exemplo 15  (Recursion) \n\n")
	underBlue.Printf(Explain.ExplainThis(15))
	exampleText.Print("\nEXEMPLO : \n")
	fmt.Print("\n Esta função fact chama a si mesma até atingir o caso base de fact(0).\n")
	exampleStyle.Print(`
	package main

	import "fmt"
	func fact(n int) int {
		if n == 0 {
			return 1
		}
		return n * fact(n-1)
	}
	func main() {
		fmt.Println(fact(7))
	`)
	underBlue.Print("\n Os fechamentos também podem ser recursivos, mas isso requer que o fechamento seja declarado com um digitado varexplicitamente antes de ser definido.\n")
	exampleStyle.Print(`
		var fib func(n int) int

		fib = func(n int) int {
			if n < 2 {
				return n
			}
			return fib(n-1) + fib(n-2)
		}
	`)
	underBlue.Print("\nComo fib foi declarado anteriormente em main, Go sabe com qual função chamar fib aqui.\n")
	exampleStyle.Print(`

		fmt.Println(fib(7))
	}
	`)
	resultText.Print("\nRESULTADO : \n")
	fmt.Println(fact(7))
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))

}
