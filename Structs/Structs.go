package Structs

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

type pessoa struct {
	nome  string
	idade int
}

func novaPessoa(nome string) *pessoa {

	p := pessoa{nome: nome}
	p.idade = 42
	return &p
}

func Structs() {
	underRed.Print("\n\n\n Exemplo 17  (Structs) \n\n")
	underBlue.Printf(Explain.ExplainThis(17))
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	package main
	
	import "fmt"
	`)
	underBlue.Print("\n O tipo Struct pessoa possui campos de nome e idade\n")
	exampleStyle.Print(`
	type pessoa struct {
		nome string
		idade  int
	}
	`)
	underBlue.Print("\n novaPessoa constrói uma nova pessoa com o nome fornecido\n")
	exampleStyle.Print(`
	
	func novaPessoa(nome string) *pessoa {

		p := pessoa{nome: nome}
		p.idade = 42
	`)
	underBlue.Print("\n Você pode retornar com segurança um ponteiro para a variável local, pois uma variável local sobreviverá ao escopo da função.\n")
	exampleStyle.Print(`
		return &p
	}
	
	func main() {
	`)
	underBlue.Print("\n Esta sintaxe cria uma nova estrutura(Struct)\n")
	exampleStyle.Print(`
		fmt.Println(pessoa{"Barbara", 20})
	`)
	underBlue.Print("\n Você pode nomear os campos ao inicializar uma estrutura\n")
	exampleStyle.Print(`
		fmt.Println(pessoa{nome: "Aline", idade: 30})
	`)
	underBlue.Print("\n Os campos omitidos terão valor zero\n")
	exampleStyle.Print(`		
		fmt.Println(pessoa{nome: "Fábio"})
	`)
	underBlue.Print("\n O prefixo & produz um ponteiro para a estrutura.\n")
	exampleStyle.Print(`		
		fmt.Println(&pessoa{nome: "Ana", idade: 40})
	`)
	underBlue.Print("\n É idiomático encapsular a criação de uma nova estrutura em funções construtoras\n")
	exampleStyle.Print(`		
		fmt.Println(newpessoa("João"))
	`)
	underBlue.Print("\n Acesse os campos de estrutura com um ponto.\n")
	exampleStyle.Print(`		
		s := pessoa{nome: "Sandro", idade: 50}
		fmt.Println(s.nome)
	`)
	underBlue.Print("\n Você também pode usar pontos com ponteiros de estrutura - os ponteiros são automaticamente desreferenciados.\n")
	exampleStyle.Print(`		
		sp := &s
		fmt.Println(sp.idade)
	`)
	underBlue.Print("\n As estruturas são mutáveis.\n")
	exampleStyle.Print(`		
		sp.idade = 51
		fmt.Println(sp.idade)
	}
	`)
	resultText.Print("\nRESULTADO : \n")
	fmt.Println(pessoa{"Barbara", 20})
	fmt.Println(pessoa{nome: "Aline", idade: 30})
	fmt.Println(pessoa{nome: "Fábio"})
	fmt.Println(&pessoa{nome: "Ana", idade: 40})
	fmt.Println(novaPessoa("João"))
	s := pessoa{nome: "Sandro", idade: 50}
	fmt.Println(s.nome)
	sp := &s
	fmt.Println(sp.idade)
	sp.idade = 51
	fmt.Println(sp.idade)
}
