package RangeOverChannels

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

func RangeOverChannels() {
	underRed.Print("\n\n\n Exemplo 31  (Range em Canais) \n\n")
	underBlue.Printf(Explain.ExplainThis(31))
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	package main
	
	import "fmt"
	
	func main() {
	`)
	underBlue.Print("\n Iremos iterar mais de 2 valores no canal queue. \n")
	exampleStyle.Print(`
		queue := make(chan string, 2)
		queue <- "Um"
		queue <- "Dois"
		close(queue)
	`)
	underBlue.Print("\n Este range itera sobre cada elemento à medida que é recebido de queue.\n")
	fmt.Print(" Como fechamos o canal acima, a iteração termina após receber os 2 elementos.\n")
	exampleStyle.Print(`	
		for elem := range queue {
			fmt.Println(elem)
		}
	}
	`)
	underBlue.Print("Este exemplo também mostrou que é possível fechar um canal não vazio, mas ainda receber os valores restantes.\n")
	resultText.Print("\nRESULTADO : \n\n")
	queue := make(chan string, 2)
	queue <- "Um"
	queue <- "Dois"
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}

}
