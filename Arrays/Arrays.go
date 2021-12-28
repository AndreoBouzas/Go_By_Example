package Arrays

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

func Arrays() {
	underRed.Print("\n\n\nExemplo 8  (Arrays): \n\n")
	underBlue.Println(Explain.ExplainThis(8))
	fmt.Print("criamos um array aque conterá exatamente 5 ints. O tipo de elementos e comprimento são parte do tipo da matriz. Por padrão, uma matriz tem valor zero, o que para ints significa 0s.")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	var a [5]int
	fmt.Println("emp:", a)
	`)
	resultText.Print("\nRESULTADO : \n")
	var a [5]int
	fmt.Print("emp:", a, "\n\n")
	//----------------------------------------------------------------------------
	underBlue.Print("Podemos definir um valor em um índice usando a array[índice] = Valor sintaxe e obter um valor com array[índice].")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	`)
	resultText.Print("\nRESULTADO : \n")
	a[4] = 100
	fmt.Print("set:", a, "\n\n")
	fmt.Print("get:", a[4], "\n\n")
	//----------------------------------------------------------------------------
	underBlue.Print("O builtin len retorna o comprimento de uma matriz. ")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	fmt.Println("len:", len(a))
	`)
	resultText.Print("\nRESULTADO : \n")
	fmt.Println("len:", len(a))
	//----------------------------------------------------------------------------
	underBlue.Print("podemos usar a seguinte sintaxe para declarar e inicializar um array em uma linha.")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	`)
	resultText.Print("\nRESULTADO : \n")
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	//----------------------------------------------------------------------------
	underBlue.Print("Os tipos de matriz são unidimensionais, mas você pode compor tipos para construir estruturas de dados multidimensionais.")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	`)
	resultText.Print("\nRESULTADO : \n")

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
