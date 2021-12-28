package Constants

import (
	"fmt"
	"math"

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

const s string = "constante"

func Constants() {
	underRed.Print("\n\n\n Exemplo 4  (Constantes): \n\n")
	underBlue.Println(Explain.ExplainThis(4))
	fmt.Print("const declara um valor constante. \n " + " Uma declaração const  pode aparecer em qualquer lugar que uma declaração var  possa. ")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	const  s  string  =  "constante"
	func  main ()  {
		fmt.Println(s)
		const  n  =  500000000
	`)
	resultText.Print("\nRESULTADO : \n")
	const n = 500000000
	fmt.Print(s, "\n\n")
	//----------------------------------------------------------------------------
	underBlue.Print("Expressões constantes realizam operações aritméticas com precisão arbitrária. ")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	const d = 3e20 / n
	fmt.Println(d)
	`)
	resultText.Print("\nRESULTADO : \n")
	const d = 3e20 / n
	fmt.Print(d, "\n\n")
	//----------------------------------------------------------------------------
	underBlue.Print("Uma constante numérica não tem tipo até que seja fornecida, como por meio de uma conversão explícita. ")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	fmt.Println(int64(d))
	`)
	resultText.Print("\nRESULTADO : \n")
	fmt.Print(int64(d), "\n\n")
	//----------------------------------------------------------------------------
	underBlue.Print("Um número pode receber um tipo usando-o em um contexto que requer um, como uma atribuição de variável ou chamada de função. Por exemplo, aqui math.Sin espera um float64.")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	fmt.Println(math.Sin(n))
	`)
	resultText.Print("\nRESULTADO : \n")
	fmt.Print(math.Sin(n), "\n\n")
}
