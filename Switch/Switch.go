package Switch

import (
	"fmt"
	"time"

	Explain "github.com/AndreoBouzas/Go_By_Example/Explain"
)

func Switch() {
	fmt.Print("Exemplo 7  (Switch): \n\n")
	fmt.Println(Explain.ExplainThis(7))

	fmt.Print("Abaixo está o tipo mais básico de estrutura Switch \n EXEMPLO : \n" + `
	
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
	` + "\n\n RESULTADO : \n\n")
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

	fmt.Print("Você também pode usar vírgulas para separar várias expressões na mesma declaração case. Usamos o default caso opcional neste exemplo também. \n EXEMPLO : \n" + `
	
	switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("É fim de semana! :) ")
    default:
        fmt.Println("É dia de semana! :( ")
    }
	` + "\n\n RESULTADO : \n\n")

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Print("É fim de semana! :) \n\n")
	default:
		fmt.Print("É dia de semana! :( \n\n")
	}

	fmt.Print("switch sem uma expressão, é uma forma alternativa de expressar a lógica if / else. Aqui também mostramos como as expressões case podem ser não constantes. \n EXEMPLO : \n" + `
	
	t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("É antes do meio-dia")
    default:
        fmt.Println("É depois do meio-dia")
    }
	` + "\n\n RESULTADO : \n\n")
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Print("É antes do meio-dia \n\n")
	default:
		fmt.Print("É depois do meio-dia \n\n")
	}

	fmt.Print("Switch compara tipos em vez de valores. Você pode usar isso para descobrir o tipo de um valor de interface. Neste exemplo, a variável t terá o tipo correspondente a sua cláusula. \n EXEMPLO : \n" + `
	
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
    whatAmI("hey")` + "\n\n RESULTADO : \n\n")
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
