package Functions

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

func plus(a int, b int) int {
	return a + b
}
func plusPlus(a, b, c int) int {
	return a + b + c
}
func vals() (int, int) {
	return 3, 7
}
func Functions() {
	underRed.Print("\n\n\n Exemplo 12  (Functions) \n\n")
	underBlue.Printf(Explain.ExplainThis(12))
	fmt.Print("\n Aqui a Função recebe dois valores Ints e soma os dois retornando um Int\n")
	fmt.Print("\n Go requer retornos explícitos, ou seja, não retornará automaticamente o valor da última expressão.\n")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	func plus(a int, b int) int {
		return a + b
	}
	`)
	resultText.Print("\nRESULTADO : \n")
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	//----------------------------------------------------------------------
	underBlue.Print("\n Quando você tem vários parâmetros consecutivos do mesmo tipo, pode omitir o nome do tipo para os parâmetros com tipos semelhantes até o parâmetro final que declara o tipo. \n")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	func plusPlus(a, b, c int) int {
		return a + b + c
	}
	`)
	resultText.Print("\nRESULTADO : \n")
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
	//----------------------------------------------------------------------
	underBlue.Print("\n Chame uma função conforme o esperado, com name(args).\n")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
	`)
	resultText.Print("\nRESULTADO : \n")
	res = plus(1, 2)
	fmt.Println("1+2 =", res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
	//----------------------------------------------------------------------
	underBlue.Print("\n Existem vários outros recursos para funções Go. Um é vários valores de retorno\n")
	fmt.Print("Go tem suporte integrado para vários valores de retorno . \n")
	fmt.Print("Esse recurso é frequentemente usado em Go idiomático, por exemplo, para retornar valores de resultado e erro de uma função.\n")
	fmt.Print("A (int, int)assinatura nesta função mostra que a função retorna 2 ints.\n")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	func vals() (int, int) {
		return 3, 7
	}
	`)
	//----------------------------------------------------------------------
	underBlue.Print("\n Aqui usamos os 2 valores de retorno diferentes da chamada com atribuição múltipla .\n")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	`)
	resultText.Print("\nRESULTADO : \n")
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	//----------------------------------------------------------------------
	underBlue.Print("\nSe você deseja apenas um subconjunto dos valores retornados, use o identificador em branco _.\n")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	_, c := vals()
	fmt.Println(c)
	`)
	resultText.Print("\nRESULTADO : \n")
	_, c := vals()
	fmt.Println(c)
	//----------------------------------------------------------------------

}
