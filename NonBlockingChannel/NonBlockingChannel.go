package NonBlockingChannel

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

func NonBlockingChannel() {
	underRed.Print("\n\n\n Exemplo 29  (Canais sem bloqueio) \n\n")
	underBlue.Printf(Explain.ExplainThis(29))
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	package main
	
	import "fmt"
	
	func main() {
		mensagens := make(chan string)
		sinais := make(chan bool)
	`)
	underBlue.Print("\n Se um valor estiver disponível no canal 'mensagens' então, select polula a variável 'msg' com o valor. \n")
	fmt.Print(" Caso contrário, ele irá imediatamente rodar o case default.\n")
	exampleStyle.Print(`
		select {
		case msg := <-mensagens:
			fmt.Println("Mensagem recebida", msg)
		default:
			fmt.Println("Nenhuma mensagen recebida!")
		}
	`)
	underBlue.Print("\n Um canal de envio sem bloqueio funciona de forma similar\n")
	fmt.Print(" Aqui 'msg' não consegue enviar para o canal 'mensagens',pois o canal não possui buffer e não há receptor \n")
	fmt.Print(" Portanto o case default é selecionado\n")
	exampleStyle.Print(`
		msg := "hi"
		select {
		case mensagens <- msg:
			fmt.Println("Enviando mensagem", msg)
		default:
			fmt.Println("Nenhuma mensagem enviada!")
		}
	`)
	underBlue.Print("\n É possível usar multiplos cases a cima do case default para implementar um select multidirecional sem bloqueio \n")
	fmt.Print(" Aqui tentamos receber sem bloqueio tento em 'mensagens' quanto em 'sinais'.\n")
	exampleStyle.Print(`
		select {
		case msg := <-mensagens:
			fmt.Println("Mensagem recebida", msg)
		case sig := <-sinais:
			fmt.Println("Sinal recebido", sig)
		default:
			fmt.Println("Sem atividade")
		}
	}
	`)
	resultText.Print("\nRESULTADO : \n\n")
	mensagens := make(chan string)
	sinais := make(chan bool)
	select {
	case msg := <-mensagens:
		fmt.Println("Mensagem recebida", msg)
	default:
		fmt.Println("Nenhuma mensagen recebida!")
	}
	msg := "hi"
	select {
	case mensagens <- msg:
		fmt.Println("Enviando mensagem", msg)
	default:
		fmt.Println("Nenhuma mensagem enviada")
	}
	select {
	case msg := <-mensagens:
		fmt.Println("Mensagem recebida", msg)
	case sig := <-sinais:
		fmt.Println("Sinal recebido", sig)
	default:
		fmt.Println("Sem atividade")
	}

}
