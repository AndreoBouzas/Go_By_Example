package ClosingChannels

import (
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

func ClosingChannels() {
	underRed.Print("\n\n\n Exemplo 30  (Fechando Canais) \n\n")
	underBlue.Printf(Explain.ExplainThis(30))
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	package main
	
	import "fmt"
	`)
	underBlue.Print("\n Neste exemplo, usaremos um canal jobs para comunicar o trabalho a ser feito da goroutine main(), para uma goroutine do trabalhador.\n")
	fmt.Print(" Quando não tivermos mais empregos para o trabalhador vamos fechar(close) o canal.\n")
	exampleStyle.Print(`
	func main() {
		jobs := make(chan int, 5)
		done := make(chan bool)
	`)
	underBlue.Print("\n Aqui está a goroutine do trabalhador. Ele recebe repetidamente de jobs com (j, more := <-jobs). \n")
	fmt.Print(" Nesta forma especial de recebimento de 2 valores, o valor de 'more' será false se jobs foi fechado e todos os valores no canal já foram recebidos.\n")
	fmt.Print(" Usamos isso para notificar 'done' quando trabalhamos em todos os nossos trabalhos.\n")
	exampleStyle.Print(`	
		go func() {
			for {
				j, more := <-jobs
				if more {
					fmt.Println("Trabalho recebido", j)
				} else {
					fmt.Println("Todos os trabalhos recebidos")
					done <- true
					return
				}
			}
		}()
	`)
	underBlue.Print("\n Isso envia 3 jobs para o trabalhador pelo canal jobs e o fecha. \n")
	exampleStyle.Print(`	
		for j := 1; j <= 3; j++ {
			jobs <- j
			fmt.Println("Enviendo trabalho", j)
		}
		close(jobs)
		fmt.Println("todos os trabalhos recebidos")
	`)
	underBlue.Print("\n Aguardamos o trabalhador usando a abordagem de sincronização que vimos anteriormente. \n")
	exampleStyle.Print(`
		<-done
	}
	`)
	resultText.Print("\nRESULTADO : \n\n")
	jobs := make(chan int, 5)
	done := make(chan bool)
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Trabalho recebido", j)
			} else {
				fmt.Println("Todos os trabalhos recebidos")
				done <- true
				return
			}
		}
	}()
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("Enviando trabalho", j)
	}
	close(jobs)
	fmt.Println("todos os trabalhos enviados")
	<-done
}
