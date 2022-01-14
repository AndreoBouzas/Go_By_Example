package Timeouts

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

func Timeouts() {
	underRed.Print("\n\n\n Exemplo 28  (TimeOuts) \n\n")
	underBlue.Printf(Explain.ExplainThis(28))
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	package main
	
	import (
		"fmt"
		"time"
	)
	
	func main() {
	`)
	underBlue.Print("\n Suponha que estamos executando uma chamada externa que retorna seu resultado em um canal 'c1' após 2s.  \n")
	fmt.Print(" Observe que o canal é armazenado em buffer, portanto, o envio na goroutine não é bloqueante.\n")
	fmt.Print(" Este é um padrão comum para evitar vazamentos de goroutine caso o canal nunca seja lido.\n")
	exampleStyle.Print(`
		c1 := make(chan string, 1)
		go func() {
			time.Sleep(2 * time.Second)
			c1 <- "Resultado 1"
		}()
	`)
	underBlue.Print("\n Neste trecho o Select implementa um Timeout\n")
	fmt.Print(" 'res := <-c1' Aguarda o resultado, e '<-time.After' aguarda um valor a ser enviado após o timeout de 1s. \n")
	fmt.Print(" Como select continua com o primeiro recebimento que está pronto, usaremos o caso de tempo limite se a operação demorar mais do que os 1s permitidos.\n")
	exampleStyle.Print(`
		select {
		case res := <-c1:
			fmt.Println(res)
		case <-time.After(1 * time.Second):
			fmt.Println("timeout 1")
		}
	`)
	underBlue.Print("\n Se permitirmos um tempo limite maior de 3s, o recebimento de 'c2' será bem-sucedido e imprimiremos o resultado. \n")
	exampleStyle.Print(`	
		c2 := make(chan string, 1)
		go func() {
			time.Sleep(2 * time.Second)
			c2 <- "Resultado 2"
		}()
		
		select {
		case res := <-c2:
			fmt.Println(res)
		case <-time.After(3 * time.Second):
			fmt.Println("timeout 2")
		}
	}
	`)
	underBlue.Print("A execução deste programa mostra o tempo limite da primeira operação e o sucesso da segunda.\n")
	resultText.Print("\nRESULTADO : \n\n")
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Resultado 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Resultado 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
