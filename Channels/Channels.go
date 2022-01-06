package Channels

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

func Channels() {
	underRed.Print("\n\n\n Exemplo 23  (Canais) \n\n")
	underBlue.Printf(Explain.ExplainThis(23))
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	package main
	
	import "fmt"
	
	func main() {
	`)
	underBlue.Print("\n Crie um novo canal com make(chan val-type). Os canais são digitados pelos valores que transmitem.\n")
	exampleStyle.Print(`		
		messages := make(chan string)
	`)
	underBlue.Print("\n Envie um valor para um canal usando a sintaxe channel <- .\n")
	fmt.Print(" Aqui enviamos 'ping' para o messages canal que fizemos acima, a partir de um novo goroutine.\n")
	exampleStyle.Print(`
		go func() { messages <- "ping" }()
	`)
	underBlue.Print("\n A sintaxe <-channel ,recebe um valor do canal. \n")
	fmt.Print(" Aqui receberemos a mensagem 'ping'  que enviamos acima e a imprimiremos.\n")
	exampleStyle.Print(`
		msg := <-messages
		fmt.Println(msg)
	}
	`)
	underBlue.Print("\n Quando executamos o programa, a mensagem 'ping' é passada com sucesso de um goroutine para outro por meio de nosso canal.\n")
	fmt.Print(" Por padrão, envia e recebe bloqueio até que o remetente e o destinatário estejam prontos.\n")
	fmt.Print(" Esta propriedade permitiu-nos aguardar no final do nosso programa pela mensagem 'ping' sem ter que usar qualquer outra sincronização.\n\n")
	resultText.Print("\nRESULTADO : \n\n")
	messages := make(chan string)
	go func() { messages <- "ping" }()
	msg := <-messages
	fmt.Println(msg)
}
