package Timers

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

func Timers() {
	underRed.Print("\n\n\n Exemplo 32  (Temporizadores) \n\n")
	underBlue.Printf(Explain.ExplainThis(32))
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	package main
	
	import (
		"fmt"
		"time"
	)
	
	func main() {
	`)
	underBlue.Print("\n Os Timers(temporizadores) representam um único evento no futuro. \n")
	fmt.Print(" Você informa ao cronômetro quanto tempo deseja esperar e ele fornece um canal que será notificado nesse momento.\n")
	fmt.Print(" Este temporizador irá esperar 2 segundos.\n")
	exampleStyle.Print(`
		timer1 := time.NewTimer(2 * time.Second)
	`)
	underBlue.Print("\n O <-timer1.C bloqueia no canal C do timer até que ele envie um valor indicando que o timer disparou.\n")
	exampleStyle.Print(`
		<-timer1.C
		fmt.Println("Temporizador 1 disparado")
	`)
	underBlue.Print("\n If you just wanted to wait, you could have used time.Sleep.\n")
	fmt.Print(" Uma razão pela qual um cronômetro pode ser útil é que você pode cancelar o cronômetro antes que ele seja acionado. Aqui está um exemplo disso.\n")
	exampleStyle.Print(`	
		timer2 := time.NewTimer(time.Second)
		go func() {
			<-timer2.C
			fmt.Println("Temporizador 2 disparado")
		}()
		
		stop2 := timer2.Stop()
		if stop2 {
			fmt.Println("Temporizador 2 parado")
		}
	`)
	underBlue.Print("\n Dê ao timer2 tempo suficiente para disparar, se é que vai disparar, para mostrar que está de fato parado. \n")
	exampleStyle.Print(`
		time.Sleep(2 * time.Second)
	}
	`)
	underBlue.Print(" O primeiro temporizador irá disparar ~2s depois de iniciarmos o programa, mas o segundo deve ser interrompido antes de ter a chance de disparar.\n")
	resultText.Print("\nRESULTADO : \n\n")
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Temporizador 1 disparado")
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Temporizador 2 disparado")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Temporizador 2 parado")
	}
	time.Sleep(2 * time.Second)

}
