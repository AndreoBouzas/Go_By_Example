package Explain

import (
	"fmt"
)

func ExplainThis(exampleNumber int) string {
	var finaltext string
	if exampleNumber == 1 {
		finaltext := fmt.Sprintln("Este programa aprenta um print visual de uma mensagem padrão!")
		return finaltext
	} else if exampleNumber == 2 {
		finaltext = fmt.Sprintln("Este programa aprenta os tipos primitivos de valores em go!")
	} else if exampleNumber == 3 {
		finaltext = fmt.Sprintln("Este programa aprenta as formas de declarar e inicializar as variáveis!")
	} else if exampleNumber == 4 {
		finaltext = fmt.Sprintln("Este programa aprenta as constantes, formas de declarar e inicializa-las!")
	} else if exampleNumber == 5 {
		finaltext = fmt.Sprintln("for é a única construção de loop de Go. Aqui estão alguns tipos básicos de forloops!")
	} else if exampleNumber == 6 {
		finaltext = fmt.Sprintln("If e Else, são condições de verificação que determinam diferentes comportamentos do código")
	}
	return finaltext
}
