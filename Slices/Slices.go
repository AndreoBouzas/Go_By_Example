package Slices

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

func Slices() {
	underRed.Print("\n\n\n Exemplo 9  (Slices) \n\n")
	underBlue.Printf(Explain.ExplainThis(9))
	fmt.Print("Ao contrário das matrizes, as fatias são digitadas apenas pelos elementos que contêm (não pelo número de elementos).")
	fmt.Print("\nCriar uma fatia vazia com comprimento diferente de zero, use o embutido make. Aqui, fazemos uma fatia de strings de comprimento 3(inicialmente com valor zero). \n")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	fmts := make([]string, 3)
	fmt.Println("emp:", s)
	`)
	resultText.Print("\nRESULTADO : \n")
	s := make([]string, 3)
	fmt.Print("emp:", s, "\n\n")
	//--------------------------------------------------------------------------------------------
	underBlue.Print("É possível definir e obter da mesma forma que os arrays.")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	`)
	resultText.Print("\nRESULTADO : \n")
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Print("get:", s[2], "\n\n")
	//----------------------------------------------------------------------------------------
	underBlue.Print("len retorna o comprimento da fatia conforme o esperado.")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	fmt.Println("len:", len(s))
	`)
	resultText.Print("\nRESULTADO : \n")
	fmt.Print("len:", len(s), "\n\n")
	//---------------------------------------------------------------------------------------
	underBlue.Println("Além dessas operações básicas, as fatias oferecem suporte a várias outras que as tornam mais ricas do que os arrays.")
	fmt.Println("Um é o embutido append, que retorna uma fatia contendo um ou mais novos valores. ")
	fmt.Print("Observe que precisamos aceitar um valor de retorno de append pois podemos obter um novo valor de fatia.")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	`)
	resultText.Print("\nRESULTADO : \n")
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Print("apd:", s, "\n\n")
	//---------------------------------------------------------------------------------------
	underBlue.Print("Slices também podem ser copiadas! Aqui nós criamos um slice vazio C com o mesmo comprimento de S, então copiamos S em C")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)
	`)
	resultText.Print("\nRESULTADO : \n")
	c := make([]string, len(s))
	copy(c, s)
	fmt.Print("cpy:", c, "\n\n")
	//----------------------------------------------------------------------------------------
	underBlue.Print("As fatias oferecem suporte a um operador “fatia” com a sintaxe slice[low:high]. Por exemplo, este recebe uma fatia dos elementos s[2], s[3]e s[4].")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	l := s[2:5]
	fmt.Println("sl1:", l)
	`)
	resultText.Print("\nRESULTADO : \n")
	l := s[2:5]
	fmt.Print("sl1:", l, "\n\n")
	//----------------------------------------------------------------------------------------
	underBlue.Print("Isso divide até (mas excluindo) s[5].")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	l = s[:5]
	fmt.Println("sl2:", l)
	`)
	resultText.Print("\nRESULTADO : \n")
	l = s[:5]
	fmt.Print("sl2:", l, "\n\n")
	//----------------------------------------------------------------------------------------
	underBlue.Print("E isso divide (e incluindo) s[2].")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	l = s[2:]
	fmt.Println("sl3:", l)
	`)
	resultText.Print("\nRESULTADO : \n")
	l = s[2:]
	fmt.Print("sl3:", l, "\n\n")
	//----------------------------------------------------------------------------------------
	underBlue.Print("Também podemos declarar e inicializar uma variável para a fatia em uma única linha.")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	`)
	resultText.Print("\nRESULTADO : \n")
	t := []string{"g", "h", "i"}
	fmt.Print("dcl:", t, "\n\n")
	//---------------------------------------------------------------------------------------
	underBlue.Print("As fatias podem ser compostas em estruturas de dados multidimensionais. O comprimento das fatias internas pode variar, ao contrário de matrizes multidimensionais.")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {	
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {	
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	`)
	resultText.Print("\nRESULTADO : \n")
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Print("2d: ", twoD, "\n\n")
	//------------------------------------------------------------------------------------------
	color.Green("\n\nObserve que, embora as fatias sejam de tipos diferentes dos arrays, elas são renderizadas de forma semelhante por fmt.Println.\n\n")
}
