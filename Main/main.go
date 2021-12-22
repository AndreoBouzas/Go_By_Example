package main

import (
	"fmt"

	Constants "github.com/AndreoBouzas/Go_By_Example/Constants"
	For "github.com/AndreoBouzas/Go_By_Example/For"
	Hello_World "github.com/AndreoBouzas/Go_By_Example/Hello_World"
	If_Else "github.com/AndreoBouzas/Go_By_Example/If_Else"
	Values "github.com/AndreoBouzas/Go_By_Example/Values"
	Variables "github.com/AndreoBouzas/Go_By_Example/Variables"
)

func main() {
	fmt.Println("Digite o NÚMERO do exemplo que gostaria de ohar! :")
	fmt.Println("Exemplo 1  (Hello World!)")
	fmt.Println("Exemplo 2  (Valores e seus tipos)")
	fmt.Println("Exemplo 3  (Variáveis e Declarações)")
	fmt.Println("Exemplo 4  (Constantes)")
	fmt.Println("Exemplo 5  (Laço For)")
	fmt.Println("Exemplo 6  (If/Else)")
	fmt.Println("")
	exampleNumber := 0
	fmt.Scan(&exampleNumber)

	if exampleNumber == 1 {
		Hello_World.HelloWorld()
	} else if exampleNumber == 2 {
		Values.Values()
	} else if exampleNumber == 3 {
		Variables.Variables()
	} else if exampleNumber == 4 {
		Constants.Constants()
	} else if exampleNumber == 5 {
		For.For()
	} else if exampleNumber == 6 {
		If_Else.If_Else()
	}

}
