package Methods

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

type retangulo struct {
	largura, altura int
}

func (r *retangulo) area() int {
	return r.largura * r.altura
}

func (r retangulo) perimetro() int {
	return 2*r.largura + 2*r.altura
}

func Methods() {
	underRed.Print("\n\n\n Exemplo 18  (Métodos) \n\n")
	underBlue.Printf(Explain.ExplainThis(18))
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	package main
	
	import "fmt"
	
	type retangulo struct {
		largura, altura int
	}
	`)
	underBlue.Print("\n Este método area possui um tipo receptor de *retangulo.\n")
	exampleStyle.Print(`
	func (r *retangulo) area() int {
		return r.largura * r.altura
	}
	`)
	underBlue.Print("\n Os métodos podem ser definidos para os tipos de ponteiro ou receptor de valor. Aqui está um exemplo de um receptor de valor.\n")
	exampleStyle.Print(`
	func (r retangulo) perimetro() int {
		return 2*r.largura + 2*r.altura
	}
	
	func main() {
		r := retangulo{largura: 20, altura: 10}
	`)
	underBlue.Print("\n Aqui chamamos os 2 métodos definidos para nossa estrutura.\n")
	exampleStyle.Print(`
		fmt.Println("Área: ", r.area())
		fmt.Println("Perímetro:", r.perimetro())
	`)
	underBlue.Print("\n Go lida automaticamente com a conversão entre valores e ponteiros para chamadas de método.\n")
	fmt.Print(" Você pode querer usar um tipo de receptor de ponteiro para evitar a cópia em chamadas de método ou para permitir que o método modifique a estrutura de recebimento.\n")
	exampleStyle.Print(`
		rp := &r
		fmt.Println("Área: ", rp.area())
		fmt.Println("Perímetro:", rp.perimetro())
	}
	`)
	resultText.Print("\nRESULTADO : \n")
	r := retangulo{largura: 20, altura: 10}
	fmt.Println("Área: ", r.area())
	fmt.Println("Perímetro:", r.perimetro())
	rp := &r
	fmt.Println("Área: ", rp.area())
	fmt.Println("Perímetro:", rp.perimetro())
}
