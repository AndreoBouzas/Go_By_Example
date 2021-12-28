package Value

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

func Values() {
	underRed.Print("Exemplo 2  (Valores e seus tipos): \n\n")
	underBlue.Println(Explain.ExplainThis(2))
	fmt.Print("Strings, Podem ser adicionadas com o +")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	fmt.Println("go" + "lang")
	`)
	resultText.Print("\nRESULTADO : \n")
	fmt.Print("go" + "lang\n\n")
	//---------------------------------------------------------------------------
	underBlue.Print("Integers e floats.")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	`)
	resultText.Print("\nRESULTADO : \n")
	fmt.Println("1+1 =", 1+1)
	fmt.Print("7.0/3.0 =", 7.0/3.0, "\n\n")
	//----------------------------------------------------------------------------
	underBlue.Print("Booleans, uso dos operadores booleanos como esperado, true/ false ")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	fmt.Println(true || false)
	fmt.Println(true && false)
	fmt.Println(!true)
	`)
	resultText.Print("\nRESULTADO : \n")
	fmt.Println(true || false)
	fmt.Println(true && false)
	fmt.Println(!true)
}
