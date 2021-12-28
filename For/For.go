package For

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

func For() {
	underRed.Print("\n\n\nExemplo 5  (Laço For): \n\n")
	underBlue.Println(Explain.ExplainThis(5))
	fmt.Print("O tipo mais básico, com uma única condição.")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	i := 1
	for i <= 3 {
		fmt.Print(i, ",")
		i = i + 1
	}
	`)
	resultText.Print("\nRESULTADO : \n")
	i := 1
	for i <= 3 {
		fmt.Print(i, ",")
		i = i + 1
	}
	fmt.Print("\n\n")
	//----------------------------------------------------------------------------
	underBlue.Print("A estrutura clássica estadoInicial/CondiçãoDoLaço/PósCondição :")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	for j := 7; j <= 9; j++ {
		fmt.Print(j, ",")
	}
	`)
	resultText.Print("\nRESULTADO : \n")
	for j := 7; j <= 9; j++ {
		fmt.Print(j, ",")
	}
	fmt.Print("\n\n")
	//----------------------------------------------------------------------------
	underBlue.Print("Um laço for sem condição, gera um loop infinito que somente é interrompido por chaves como:  \n break ( para o laço ) ! \n return ( para retornar do laço )")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	for {
		fmt.Print("loop ,")
		break
	}
	`)
	resultText.Print("\nRESULTADO : \n")
	for {
		fmt.Print("loop , ")
		break
	}
	fmt.Print("\n\n")
	//----------------------------------------------------------------------------
	underBlue.Print("Também pode ser usado o continue que leva para a próxima iteração do loop : ")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Print(n ,", ")
	}
	`)
	resultText.Print("\nRESULTADO : \n")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Print(n, ", ")
	}
	fmt.Println("")
}
