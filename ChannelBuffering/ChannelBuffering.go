package ChannelBuffering

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

func ChannelBuffering() {
	underRed.Print("\n\n\n Exemplo 24  (Buffer de canal) \n\n")
	underBlue.Printf(Explain.ExplainThis(24))
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	package main
	
	import "fmt"
	
	func main() {
	`)
	underBlue.Print("\n Aqui temos um canal make de strings armazenando em buffer até 2 valores.\n")
	exampleStyle.Print(`
		messages := make(chan string, 2)
	`)
	underBlue.Print("\n Como esse canal é armazenado em buffer, podemos enviar esses valores para o canal sem um recebimento simultâneo correspondente.\n")
	exampleStyle.Print(`
		messages <- "buffered"
		messages <- "channel"
	`)
	underBlue.Print("\n Posteriormente, podemos receber esses dois valores como de costume.\n")
	exampleStyle.Print(`
		fmt.Println(<-messages)
		fmt.Println(<-messages)
	}
	`)
	resultText.Print("\nRESULTADO : \n\n")
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
