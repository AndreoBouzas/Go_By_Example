package Goroutines

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

func f(De string) {
	for i := 0; i < 6; i++ {
		fmt.Println(De, ":", i)
	}
}

func Goroutines() {
	underRed.Print("\n\n\n Exemplo 22  (Goroutines) \n\n")
	underBlue.Printf(Explain.ExplainThis(22))
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	package main
	import (
		"fmt"
		"time"
	)
	
	func f(De string) {
		for i := 0; i < 6; i++ {
			fmt.Println(De, ":", i)
		}
	}
	
	func main() {
	`)
	underBlue.Print("\n Suponha que temos uma chamada de função f(s). É assim que chamaríamos isso da maneira usual, executando-o de forma síncrona.\n")
	exampleStyle.Print(`
		f("Diretamente")
	`)
	underBlue.Print("\n Para invocar esta função em uma goroutine, use go f(s). Esta nova goroutine será executada simultaneamente com a que está chamando.\n")
	exampleStyle.Print(`
		go f("Goroutine")
	`)
	underBlue.Print("\n Você também pode iniciar uma goroutine para uma chamada de função anônima.\n")
	exampleStyle.Print(`
		go func(msg string) {
			fmt.Println(msg)
		}("Indo")
	`)
	underBlue.Print("\n Nossas duas chamadas de função estão sendo executadas de forma assíncrona em goroutines separadas agora.\n")
	fmt.Print(" Espere até que eles terminem (para uma abordagem mais robusta, use um WaitGroup ).\n")
	exampleStyle.Print(`
		time.Sleep(time.Second)
		fmt.Println("Concluído")
	}
	`)
	underBlue.Print("\n Quando executamos este programa, vemos a saída da chamada de bloqueio primeiro, depois a saída das duas goroutines.\n")
	fmt.Print(" A saída dos goroutines pode ser intercalada, porque os goroutines estão sendo executados simultaneamente pelo tempo de execução Go..\n")
	resultText.Print("\nRESULTADO : \n")
	f("Diretamente")
	go f("Goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("Indo")
	time.Sleep(time.Second)
	fmt.Println("Concluído")

}
