package Maps

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

func Maps() {
	underRed.Print("\n\n\n Exemplo 10  (Maps) \n\n")
	underBlue.Printf(Explain.ExplainThis(10))
	fmt.Print("\nPara criar um mapa vazio, use o builtin make: make(map[key-type]val-type). \n")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	m := make(map[string]int)
	`)
	underBlue.Print("\nDefina pares de chave / valor usando uma name[key] = val sintaxe típica .\n")
	fmt.Print("\n\n")
	exampleStyle.Print(`
	m["k1"] = 7
	m["k2"] = 13
	`)
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	underBlue.Print("\nImprimir um mapa com, por exemplo fmt.Println, mostrará todos os seus pares de chave / valor.\n")
	exampleText.Print("\n\n")
	exampleStyle.Print(`
	fmt.Println("map:", m)
	`)
	resultText.Print("\nRESULTADO : \n")
	fmt.Println("map:", m)
	//----------------------------------------------------------------------------
	underBlue.Print("\nObtenha um valor para uma chave com name[key].\n")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	`)
	resultText.Print("\nRESULTADO : \n")
	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	//----------------------------------------------------------------------------
	underBlue.Print("\nO builtin len retorna o número de pares de chave / valor quando chamado em um mapa.\n")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	fmt.Println("len:", len(m))
	`)
	resultText.Print("\nRESULTADO : \n")
	fmt.Println("len:", len(m))
	//----------------------------------------------------------------------------
	underBlue.Print("\nO integrado delete remove pares de chave/valor de um mapa.\n")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	delete(m, "k2")
	fmt.Println("map:", m)
	`)
	resultText.Print("\nRESULTADO : \n")
	delete(m, "k2")
	fmt.Println("map:", m)
	//-----------------------------------------------------------------------------
	underBlue.Print("\nO segundo valor de retorno opcional ao obter um valor de um mapa indica se a chave estava presente no mapa.\n")
	fmt.Print(`Isso pode ser usado para eliminar a ambigüidade entre chaves ausentes e chaves com valores zero como 0 ou "" . `)
	fmt.Print(`Aqui, não precisamos do valor em si, então o ignoramos com o identificador em branco _ `)
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	`)
	resultText.Print("\nRESULTADO : \n")
	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	//----------------------------------------------------------------------------
	underBlue.Print("\nVocê também pode declarar e inicializar um novo mapa na mesma linha com esta sintaxe.\n")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
	`)
	resultText.Print("\nRESULTADO : \n")
	underBlue.Print("\nObserve que os mapas aparecem no formulário map[k:v k:v] quando impressos com fmt.Println.\n")

}
