package Explain

import (
	"fmt"
)

func ExplainThis(exampleNumber int) string {
	var finaltext string
	switch exampleNumber {
	case 1:
		finaltext = fmt.Sprintln("Este programa aprenta um print visual de uma mensagem padrão!")
	case 2:
		finaltext = fmt.Sprintln("Este programa aprenta os tipos primitivos de valores em go!")
	case 3:
		finaltext = fmt.Sprintln("Este programa aprenta as formas de declarar e inicializar as variáveis!")
	case 4:
		finaltext = fmt.Sprintln("Este programa aprenta as constantes, formas de declarar e inicializa-las!")
	case 5:
		finaltext = fmt.Sprintln("for é a única construção de loop de Go. Aqui estão alguns tipos básicos de forloops!")
	case 6:
		finaltext = fmt.Sprintln("If e Else, são condições de verificação que determinam diferentes comportamentos do código")
	case 7:
		finaltext = fmt.Sprintln("Declarações Switch expressam condicionais voltadas a diversos casos, apresentam funcionalidade similar ao If/Else!")
	case 8:
		finaltext = fmt.Sprintln("Em Go, uma matriz(Array) é uma sequência numerada de elementos de um comprimento específico.")
	case 9:
		finaltext = fmt.Sprintln("Slices são um tipo de dados chave em Go, fornecendo uma interface mais poderosa para sequências do que arrays.")
	}
	return finaltext
}
