package Errors

import (
	"errors"
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

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func Errors() {
	underRed.Print("\n\n\n Exemplo 21  (Erros) \n\n")
	underBlue.Printf(Explain.ExplainThis(21))
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	package main
	import (
		"errors"
		"fmt"
	)
	`)
	underBlue.Print("\n Por convenção, os erros são o último valor de retorno e têm tipo error, uma interface integrada.\n")
	exampleStyle.Print(`
	func f1(arg int) (int, error) {
		if arg == 42 {
	`)
	underBlue.Print("\n errors.New constrói um valor error básico com a mensagem de erro fornecida.\n")
	exampleStyle.Print(`
			return -1, errors.New("can't work with 42")
		}
	`)
	underBlue.Print("\n Um valor nil na posição de erro indica que não houve erro.\n")
	exampleStyle.Print(`
		return arg + 3, nil
	}
	`)
	underBlue.Print("\n É possível usar tipos personalizados como errors implementando o método Error() neles. \n")
	fmt.Print(" Aqui está uma variante do exemplo acima que usa um tipo personalizado para representar explicitamente um erro de argumento.\n")
	exampleStyle.Print(`
	type argError struct {
		arg  int
		prob string
	}
	
	func (e *argError) Error() string {
		return fmt.Sprintf("%D - %S", e.arg, e.prob)
	}
	
	func f2(arg int) (int, error) {
		if arg == 42 {
	`)
	underBlue.Print("\n Nesse caso, usamos a sintaxe &argError para construir uma nova estrutura, fornecendo valores para os dois campos arge prob.\n")
	exampleStyle.Print(`
			return -1, &argError{arg, "can't work with it"}
		}
		
		return arg + 3, nil
	}
	
	func main() {
	`)
	underBlue.Print("\n Os dois loops abaixo testam cada uma de nossas funções de retorno de erro.\n")
	fmt.Print(" Observe que o uso de uma verificação de erro inline na linha if é um idioma comum no código Go. \n")
	exampleStyle.Print(`
		for _, i := range []int{7, 42} {
			if r, e := f1(i); e != nil {
				fmt.Println("f1 failed:", e)
			} else {
				fmt.Println("f1 worked:", r)
			}
		}
		
		
		for _, i := range []int{7, 42} {
			if r, e := f2(i); e != nil {
				fmt.Println("f2 failed:", e)
			} else {
				fmt.Println("f2 worked:", r)
			}
		}
	`)
	underBlue.Print("\n Se desejar usar programaticamente os dados em um erro personalizado, você precisará obter o erro como uma instância do tipo de erro personalizado por meio da asserção de tipo.\n")
	exampleStyle.Print(`
		_, e := f2(42)
		if ae, ok := e.(*argError); ok {
			fmt.Println(ae.arg)
			fmt.Println(ae.prob)
		}
	}
	`)
	resultText.Print("\nRESULTADO : \n")

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
