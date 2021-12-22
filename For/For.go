package For

import (
	"fmt"

	Explain "github.com/AndreoBouzas/Go_By_Example/Explain"
)

func For() {
	fmt.Print("Exemplo 5  (Laço For): \n\n")
	fmt.Println(Explain.ExplainThis(5))
	fmt.Print("O tipo mais básico, com uma única condição. \n EXEMPLO : \n" + `
	
	i := 1
    for i <= 3 {
        fmt.Print(i, ",")
        i = i + 1
    }
	` + "\n\n RESULTADO : \n\n")
	i := 1
	for i <= 3 {
		fmt.Print(i, ",")
		i = i + 1
	}
	fmt.Print("\n\n")

	fmt.Print("A estrutura clássica estadoInicial/CondiçãoDoLaço/PósCondição : \n EXEMPLO : \n" + `
	
	for j := 7; j <= 9; j++ {
        fmt.Print(j, ",")
    }
	` + "\n\n RESULTADO : \n\n")
	for j := 7; j <= 9; j++ {
		fmt.Print(j, ",")
	}
	fmt.Print("\n\n")

	fmt.Print("Um laço for sem condição, gera um loop infinito que somente é interrompido por chaves como:  \n break ( para o laço ) ! \n return ( para retornar do laço )\n EXEMPLO : \n" + `
	
	
	for {
        fmt.Print("loop ,")
        break
    }
	` + "\n\n RESULTADO : \n\n")
	for {
		fmt.Print("loop , ")
		break
	}
	fmt.Print("\n\n")

	fmt.Print("Também pode ser usado o continue que leva para a próxima iteração do loop : \n EXEMPLO : \n" + `
	
	for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Print(n ,", ")
    }
	` + "\n\n RESULTADO : \n\n")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Print(n, ", ")
	}
	fmt.Println("")
}
