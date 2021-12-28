package Switch

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

func Switch() {
	underRed.Print("\n\n\nExemplo 7  (Switch): \n\n")
	underBlue.Println(Explain.ExplainThis(7))

	fmt.Print("Abaixo está o tipo mais básico de estrutura Switch ")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	i := 2
	fmt.Print("Escreva ", i, " como ")
	switch i {
		case 1:
			fmt.Println("Um")
		case 2:
			fmt.Println("Dois")
		case 3:
			fmt.Println("Três")
	}
	`)
	resultText.Print("\nRESULTADO : \n")
	i := 2
	fmt.Print("Escreva ", i, " como \n")
	switch i {
	case 1:
		fmt.Print("Um \n\n")
	case 2:
		fmt.Print("Dois \n\n")
	case 3:
		fmt.Print("Três \n\n")
	}
	//----------------------------------------------------------------------------
	underBlue.Print("Você também pode usar vírgulas para separar várias expressões na mesma declaração case. Usamos o default caso opcional neste exemplo também.")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("É fim de semana! :) ")
		default:
			fmt.Println("É dia de semana! :( ")
	}
	`)
	resultText.Print("\nRESULTADO : \n")
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Print("É fim de semana! :) \n\n")
	default:
		fmt.Print("É dia de semana! :( \n\n")
	}
	//----------------------------------------------------------------------------
	underBlue.Print("switch sem uma expressão, é uma forma alternativa de expressar a lógica if / else. Aqui também mostramos como as expressões case podem ser não constantes.")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("É antes do meio-dia")
		default:
			fmt.Println("É depois do meio-dia")
	}
	`)
	resultText.Print("\nRESULTADO : \n")
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Print("É antes do meio-dia \n\n")
	default:
		fmt.Print("É depois do meio-dia \n\n")
	}
	//----------------------------------------------------------------------------
	underBlue.Print("Switch compara tipos em vez de valores. Você pode usar isso para descobrir o tipo de um valor de interface. Neste exemplo, a variável t terá o tipo correspondente a sua cláusula.")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
			case bool:
				fmt.Println("Eu sou um bool")
			case int:
				fmt.Println("eu sou um Int")
			default:
				fmt.Print("Não conheço este tipo \n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	`)
	resultText.Print("\nRESULTADO : \n")
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Print("Eu sou um bool \n\n")
		case int:
			fmt.Print("eu sou um Int \n\n")
		default:
			fmt.Printf("Não conheço este tipo %T \n \n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
