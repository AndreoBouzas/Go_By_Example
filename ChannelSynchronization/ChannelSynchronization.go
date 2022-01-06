package ChannelSynchronization

import (
	"fmt"
	"time"

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

func trabalhador(concluido chan bool) {
	fmt.Print("Trabalhando...")
	time.Sleep(time.Second * 2)
	fmt.Println("concluido")
	concluido <- true
}

func ChannelSynchronization() {
	underRed.Print("\n\n\n Exemplo 25  (Sincronização de canais) \n\n")
	underBlue.Printf(Explain.ExplainThis(25))
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	package main
	
	import (
		"fmt"
		"time"
	)
	`)
	underBlue.Print("\n Esta é a função que executaremos em uma goroutine. O canal concluido será usado para notificar outra goroutine de que o trabalho desta função foi concluído.\n")
	exampleStyle.Print(`
	func trabalhador(concluido chan bool) {
		fmt.Print("Trabalhando...")
		time.Sleep(time.Second * 2)
		fmt.Println("concluido")
	`)
	underBlue.Print("\n Envia um valor para avisar que terminou.\n")
	exampleStyle.Print(`
		concluido <- true
	}
	
	func main() {
	`)
	underBlue.Print("\n Inicie uma goroutine de trabalho, dando a ela o canal para notificar.\n")
	exampleStyle.Print(`
		concluido := make(chan bool, 1)
		go trabalhador(concluido)
	`)
	underBlue.Print("\n Bloquea até recebermos uma notificação do trabalhador do canal.\n")
	exampleStyle.Print(`
		<-concluido
	}
	`)
	underBlue.Print("\n Se você remover a linha '<- concluido' deste programa, o programa 'trabalhador', seria encerrado antes mesmo de ser iniciado.\n")
	resultText.Print("\nRESULTADO : \n\n")
	concluido := make(chan bool, 1)
	go trabalhador(concluido)
	<-concluido
}
