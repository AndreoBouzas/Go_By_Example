package ChannelDirections

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

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func ChannelDirections() {
	underRed.Print("\n\n\n Exemplo 26  (Canais Direcionáis) \n\n")
	underBlue.Printf(Explain.ExplainThis(26))
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	package main
	
	import "fmt"
	`)
	underBlue.Print("\n A função 'ping' aceita apenas um canal para envio de valores. \n")
	fmt.Print(" Seria um erro em tempo de compilação tentar receber neste canal.\n")
	exampleStyle.Print(`
	func ping(pings chan<- string, msg string) {
		pings <- msg
	}

	`)
	underBlue.Print("\n A função 'pong' aceita um canal para receber (pings) e um segundo para Mandar (pongs).\n")
	exampleStyle.Print(`
	func pong(pings <-chan string, pongs chan<- string) {
		msg := <-pings
		pongs <- msg
	}
	
	func main() {
		pings := make(chan string, 1)
		pongs := make(chan string, 1)
		ping(pings, "mensagem passada")
		pong(pings, pongs)
		fmt.Println(<-pongs)
	}
	`)
	resultText.Print("\nRESULTADO : \n\n")
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "mensagem passada")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}
