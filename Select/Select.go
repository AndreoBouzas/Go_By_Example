package Select

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

func Select() {
	underRed.Print("\n\n\n Exemplo 27  (Select) \n\n")
	underBlue.Printf(Explain.ExplainThis(27))
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	package main
	
	import (
		"fmt"
		"time"
	)
	
	func main() {
	`)
	underBlue.Print("\n Nesse exemplo usaremos Select para dois canais \n")
	exampleStyle.Print(`
		c1 := make(chan string)
		c2 := make(chan string)
	`)
	underBlue.Print("\n Cada canal receberá um valor após algum tempo, para simular, por exemplo, o bloqueio de operações RPC em execução em goroutines simultâneas.\n")
	exampleStyle.Print(`
		go func() {
			time.Sleep(1 * time.Second)
			c1 <- "Um"
		}()
		
		go func() {
			time.Sleep(2 * time.Second)
			c2 <- "Dois"
		}()
	`)
	underBlue.Print("\n Usaremos select para aguardar ambos os valores simultaneamente, imprimindo cada um conforme chegar.\n")
	exampleStyle.Print(`
		for i := 0; i < 2; i++ {
			select {
			case msg1 := <-c1:
				fmt.Println("recebido", msg1)
			case msg2 := <-c2:
				fmt.Println("recebido", msg2)
			}
		}
	}
	`)
	resultText.Print("\nRESULTADO : \n\n")
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "Um"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Dois"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("recebido", msg1)
		case msg2 := <-c2:
			fmt.Println("recebido", msg2)
		}
	}
	underBlue.Print("\n Recebemos os valores 'Um' e depois 'Dois' conforme o esperado.\n")
	underBlue.Print("\n Observe que o tempo total de execução é de apenas cerca de 2 segundos, já que 1 e 2 segundos são Sleeps executados simultaneamente..\n")

}
