package Variables

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

func Variables() {
	underRed.Print("\n\n\nExemplo 3  (Variáveis e Declarações): \n\n")
	underBlue.Println(Explain.ExplainThis(3))
	fmt.Print("Var declara 1 ou mais variáveis ")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`v
	var a = "Andreo"
	fmt.Println(a)
	`)
	resultText.Print("\nRESULTADO : \n")
	var a = "Andreo"
	fmt.Print(a, "\n\n")
	//----------------------------------------------------------------------------
	underBlue.Print("Declarando multiplas vaiáveis com uso do var:")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	var b, c int = 1, 2
	fmt.Println(b, c)
	`)
	resultText.Print("\nRESULTADO : \n")
	var b, c int = 1, 2
	fmt.Print(b, c, "\n\n")
	//----------------------------------------------------------------------------
	underBlue.Print("Go associa um tipo a uma variável, basedo no valor atribuído a ela. Caso o tipo da mesma não tenha sido declarado! \n\n")
	fmt.Print("Declaração de variável do tipo booleano: ")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	var d = true
	fmt.Println(d)
	`)
	resultText.Print("\nRESULTADO : \n")
	var d = true
	fmt.Print(d, "\n\n")
	//----------------------------------------------------------------------------
	underBlue.Print("Variáveis declaradas sem valor de inicialização, recebem por padrão o zero do tipo em questão:")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	var e int
	fmt.Println(e)
	`)
	resultText.Print("\nRESULTADO : \n")
	var e int
	fmt.Print(e, "\n\n")
	//----------------------------------------------------------------------------
	underBlue.Print("O atalho de inicialização e declaração de variáveis é := ")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	f := "Maçã"
	fmt.Println(f)
	`)
	resultText.Print("\nRESULTADO : \n")
	f := "Maçã"
	fmt.Print(f, "\n\n")

}
