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
	case 19:
		finaltext = fmt.Sprintln("As interfaces são coleções nomeadas de assinaturas de método.")
	case 20:
		finaltext = fmt.Sprintln("Go oferece suporte à incorporação de structs e interfaces para expressar uma composição de tipos mais uniforme .")
	case 21:
		finaltext = fmt.Sprintln("Em Go, é idiomático comunicar erros por meio de um valor de retorno separado e explícito. Isso contrasta com as exceções usadas em linguagens como Java e Ruby e o único resultado / valor de erro sobrecarregado às vezes usado na abordagem de C. Go torna fácil ver quais funções retornam erros e manipulá-los usando as mesmas construções de linguagem empregadas para qualquer outra, tarefas sem erros.")
	case 22:
		finaltext = fmt.Sprintln("Uma goroutine é um thread de execução leve.")
	case 23:
		finaltext = fmt.Sprintln("Os canais(Channels) são os tubos que conectam goroutines simultâneas. Você pode enviar valores para canais de uma goroutine e receber esses valores em outra goroutine.")
	case 24:
		finaltext = fmt.Sprintln("Por padrão, os canais não têm buffer, o que significa que eles só aceitarão enviar( chan <-) se houver um recptor( <- chan) correspondente pronto para receber o valor enviado. Os canais com buffer aceitam um número limitado de valores sem um receptor correspondente para esses valores.")
	case 25:
		finaltext = fmt.Sprintln("Podemos usar canais para sincronizar a execução entre goroutines. Aqui está um exemplo de como usar um recebimento de bloqueio para esperar que um goroutine termine. Ao esperar que vários goroutines terminem, você pode preferir usar um WaitGroup .")
	case 26:
		finaltext = fmt.Sprintln("Ao usar canais como parâmetros de função, você pode especificar se um canal destina-se apenas a enviar ou receber valores. Essa especificidade aumenta a segurança de tipo do programa.")
	case 27:
		finaltext = fmt.Sprintln(" ")
	case 28:
		finaltext = fmt.Sprintln(" ")
	case 29:
		finaltext = fmt.Sprintln(" ")

	}
	return finaltext
}
