package If_Else

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

func If_Else() {
	underRed.Print("\n\n\n Exemplo 6  (If/Else): \n\n")
	underBlue.Println(Explain.ExplainThis(6))
	fmt.Print("Abaixo está o tipo mais básico de estrutura If Else!")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	if 7%2 == 0 {
		fmt.Println("7 é Par")
	} else {
		fmt.Println("7 é Impar")
	}
	`)
	resultText.Print("\nRESULTADO : \n")
	if 7%2 == 0 {
		fmt.Print("7 é Par", "\n\n")
	} else {
		fmt.Print("7 é Impar", "\n\n")
	}
	//----------------------------------------------------------------------------
	underBlue.Print("Também é possível declarar um If sem Else!")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	if 8%4 == 0 {
		fmt.Println("8 é divisível por 4")
	}
	`)
	resultText.Print("\nRESULTADO : \n")
	if 8%4 == 0 {
		fmt.Print("8 é divisível por 4", "\n\n")
	}
	//-----------------------------------------------------------------------------
	underBlue.Print("Uma instrução pode preceder as condicionais; quaisquer variáveis ​​declaradas nesta instrução estão disponíveis em todos os ramos.!")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	if num := 9; num < 0 {
		fmt.Println(num, "é negativo")
	} else if num < 10 {
		fmt.Println(num, "possui 1 dígito")
	} else {
		fmt.Println(num, "possui multiplos dígitos")
	}
	`)
	resultText.Print("\nRESULTADO : \n")
	if num := 9; num < 0 {
		fmt.Print(num, " é negativo", "\n\n")
	} else if num < 10 {
		fmt.Print(num, " possui 1 dígito", "\n\n")
	} else {
		fmt.Print(num, " possui multiplos dígitos", "\n\n")
	}
	underBlue.Print("Observe que você não precisa de parênteses em torno das condições em Go! \n")

}
