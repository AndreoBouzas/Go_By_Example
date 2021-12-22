package Explain

import (
	"fmt"
)

func ExplainThis(exampleNumber int) string {
	var finaltext string
	if exampleNumber == 1 {
		finaltext := fmt.Sprintln("Este programa aprenta um print visual de uma mensagem padrão!")
		return finaltext
	}
	return finaltext
}
