package Embedding

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

type base struct {
	num int
}

func (b base) descreve() string {
	return fmt.Sprintf("base com num=%v", b.num)
}

type container struct {
	base
	str string
}

func Embedding() {
	underRed.Print("\n\n\n Exemplo 20  (incorporação) \n\n")
	underBlue.Printf(Explain.ExplainThis(20))
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	package main
	
	import "fmt"
	
	type base struct {
		num int
	}
	
	func (b base) descreve() string {
		return fmt.Sprintf("base com num=%V", b.num)
	}
	`)
	underBlue.Print("\n A Struct container incorpora a Structu base. Uma incorporação parece um campo sem nome.\n")
	exampleStyle.Print(`
	type container struct {
		base
		str string
	}
	
	func main() {
	`)
	underBlue.Print("\n Ao criar Structs com literais, temos que inicializar o embedding explicitamente; aqui, o tipo incorporado serve como o nome do campo.\n")
	exampleStyle.Print(`
		co := container{
			base: base{
				num: 1,
			},
			str: "algum nome",
		}
	`)
	underBlue.Print("\n Podemos acessar os campos da base diretamente no co, por exemplo co.num.\n")
	exampleStyle.Print(`
		fmt.Printf("co={num: %V, str: %V}\n", co.num, co.str)
	`)
	underBlue.Print("\n Como alternativa, podemos soletrar o caminho completo usando o nome do tipo incorporado.\n")
	exampleStyle.Print(`
		fmt.Println("também num:", co.base.num)
	`)
	underBlue.Print("\n Já que container incorpora base, os métodos de base também se tornam métodos de a container. Aqui, invocamos um método que foi incorporado base diretamente co.\n")
	exampleStyle.Print(`
		fmt.Println("descreve:", co.descreve())
		
		type descrever interface {
			descreve() string
		}
	`)
	underBlue.Print("\n A incorporação de estruturas com métodos pode ser usada para conceder implementações de interface a outras estruturas.\n")
	fmt.Print(" Aqui, vemos que container agora implementa a interface descrever porque está embutida base.\n")
	exampleStyle.Print(`
		var d descrever = co
		fmt.Println("descrever:", d.descreve())
	}
	`)
	resultText.Print("\nRESULTADO : \n")
	co := container{
		base: base{
			num: 1,
		},
		str: "algum nome",
	}
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("também num:", co.base.num)
	fmt.Println("descreve:", co.descreve())
	type descrever interface {
		descreve() string
	}
	var d descrever = co
	fmt.Println("descrever:", d.descreve())

}
