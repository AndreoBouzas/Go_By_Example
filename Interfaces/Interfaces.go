package Interfaces

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

type geometria interface {
	area() float64
	perim() float64
}
type retangulo struct {
	largura, altura float64
}
type circulo struct {
	raio float64
}

func (r retangulo) area() float64 {
	return r.largura * r.altura
}

func (r retangulo) perim() float64 {
	return 2*r.largura + 2*r.altura
}
func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

func (c circulo) perim() float64 {
	return 2 * math.Pi * c.raio
}
func medir(g geometria) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func Interfaces() {
	underRed.Print("\n\n\n Exemplo 19  (Interfaces) \n\n")
	underBlue.Printf(Explain.ExplainThis(19))
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	package main
	import (
		"fmt"
		"math"
	)
	`)
	underBlue.Print("\n Aqui está uma interface básica para formas geométricas.\n")
	exampleStyle.Print(`
	type geometria interface {
		area() float64
		perim() float64
	}
	`)
	underBlue.Print("\n Para nosso exemplo, implementaremos essa interface em tipos retangulo e circulo.\n")
	exampleStyle.Print(`	
	type retangulo struct {
		largura, altura float64
	}
	
	type circulo struct {
		raio float64
	}
	`)
	underBlue.Print("\n Para implementar uma interface em Go, precisamos apenas implementar todos os métodos da interface. Aqui, implementamos geometria em retangulos.\n")
	exampleStyle.Print(`
	func (r retangulo) area() float64 {
		return r.largura * r.altura
	}
	
	func (r retangulo) perim() float64 {
		return 2*r.largura + 2*r.altura
	}
	`)
	underBlue.Print("\n A implementação de circulos.\n")
	exampleStyle.Print(`	
	func (c circulo) area() float64 {
		return math.Pi * c.raio * c.raio
	}
	
	func (c circulo) perim() float64 {
		return 2 * math.Pi * c.raio
	}
	`)
	underBlue.Print("\n Se uma variável tem um tipo de interface, podemos chamar métodos que estão na interface nomeada. Aqui está uma função medir genérica aproveitando isso para funcionar em qualquer geometria.\n")
	exampleStyle.Print(`
	func medir(g geometria) {
		fmt.Println(g)
		fmt.Println(g.area())
		fmt.Println(g.perim())
	}
	
	func main() {
		r := retangulo{largura: 3, altura: 4}
		c := circulo{raio: 5}
	`)
	underBlue.Print("\n Os tipos de estrutura circulo e retangulo ambos implementam a interface geometria para que possamos usar instâncias dessas estruturas como argumentos para medir.\n")
	exampleStyle.Print(`
		medir(r)
		medir(c)
	}
	`)
	resultText.Print("\nRESULTADO : \n")
	r := retangulo{largura: 3, altura: 4}
	c := circulo{raio: 5}
	medir(r)
	medir(c)
}
