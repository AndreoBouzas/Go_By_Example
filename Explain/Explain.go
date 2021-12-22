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
	}
	return finaltext
}
