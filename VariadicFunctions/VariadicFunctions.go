package VariadicFunctions

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

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func VariadicFunctions() {
	underRed.Print("\n\n\n Exemplo 13  (Função Variável) \n\n")
	underBlue.Printf(Explain.ExplainThis(13))
	fmt.Print("\n Aqui está uma função que terá um número arbitrário de ints como argumentos.\n")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	func sum(nums ...int) {
		fmt.Print(nums, " ")
		total := 0
		for _, num := range nums {
			total += num
		}
		fmt.Println(total)
	}
	`)
	//------------------------------------------------------------------------
	underBlue.Print("\n As funções variáveis ​​podem ser chamadas da maneira usual com argumentos individuais.\n")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	sum(1, 2)
	sum(1, 2, 3)
	`)
	resultText.Print("\nRESULTADO : \n")
	sum(1, 2)
	sum(1, 2, 3)
	//------------------------------------------------------------------------
	underBlue.Print("\nSe você já tem vários args em um Slice, aplique-os a uma função variável usando func(slice...)assim.\n")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	nums := []int{1, 2, 3, 4}
	sum(nums...)
	`)
	resultText.Print("\nRESULTADO : \n")
	nums := []int{1, 2, 3, 4}
	sum(nums...)

}
