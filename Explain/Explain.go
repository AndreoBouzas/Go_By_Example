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
	case 10:
		finaltext = fmt.Sprintln("Mapas são o tipo de dados associativo embutido do Go (às vezes chamados de hashes ou dicts em outras línguas).")
	case 11:
		finaltext = fmt.Sprintln("Range itera sobre elementos em uma variedade de estruturas de dados. Vamos ver como usar range em algumas das estruturas de dados que já aprendemos.")
	case 12:
		finaltext = fmt.Sprintln("As funções são centrais no Go. Aprenderemos sobre as funções com alguns exemplos diferentes.")
	case 13:
		finaltext = fmt.Sprintln("As funções variáveis podem ser chamadas com qualquer número de argumentos finais. Por exemplo,fmt.Printl né uma função variável comum.")
	case 14:
		finaltext = fmt.Sprintln("Go oferece suporte a funções anônimas , que podem formar encerramentos(Closures) . As funções anônimas são úteis quando você deseja definir uma função embutida sem precisar nomeá-la.")
	case 15:
		finaltext = fmt.Sprintln("Go oferece suporte a funções recursivas . Aqui está um exemplo clássico.")
	case 16:
		finaltext = fmt.Sprintln("Go oferece suporte a ponteiros , permitindo que você passe referências a valores e registros dentro de seu programa.")
	case 17:
		finaltext = fmt.Sprintln("As estruturas de Go são coleções digitadas de campos. Eles são úteis para agrupar dados para formar registros.")
	case 18:
		finaltext = fmt.Sprintln("Go oferece suporte a métodos definidos em tipos de estrutura.")
	}
	return finaltext
}
