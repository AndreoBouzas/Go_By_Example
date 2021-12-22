package If_Else

import (
	"fmt"

	Explain "github.com/AndreoBouzas/Go_By_Example/Explain"
)

func If_Else() {
	fmt.Print("Exemplo 6  (If/Else): \n\n")
	fmt.Println(Explain.ExplainThis(6))

	fmt.Print("Abaixo está o tipo mais básico de estrutura If Else! \n EXEMPLO : \n" + `
	
	if 7%2 == 0 {
        fmt.Println("7 é Par")
    } else {
        fmt.Println("7 é Impar")
    }
	` + "\n\n RESULTADO : \n\n")
	if 7%2 == 0 {
		fmt.Print("7 é Par", "\n\n")
	} else {
		fmt.Print("7 é Impar", "\n\n")
	}

	fmt.Print("Também é possível declarar um If sem Else! \n EXEMPLO : \n" + `
	
	if 8%4 == 0 {
        fmt.Println("8 é divisível por 4")
    }
	` + "\n\n RESULTADO : \n\n")
	if 8%4 == 0 {
		fmt.Print("8 é divisível por 4", "\n\n")
	}

	fmt.Print("Uma instrução pode preceder as condicionais; quaisquer variáveis ​​declaradas nesta instrução estão disponíveis em todos os ramos.! \n EXEMPLO : \n" + `
	
	if num := 9; num < 0 {
        fmt.Println(num, "é negativo")
    } else if num < 10 {
        fmt.Println(num, "possui 1 dígito")
    } else {
        fmt.Println(num, "possui multiplos dígitos")
    }
	` + "\n\n RESULTADO : \n\n")
	if num := 9; num < 0 {
		fmt.Print(num, " é negativo", "\n\n")
	} else if num < 10 {
		fmt.Print(num, " possui 1 dígito", "\n\n")
	} else {
		fmt.Print(num, " possui multiplos dígitos", "\n\n")
	}
	fmt.Print("Observe que você não precisa de parênteses em torno das condições em Go! \n")

}
