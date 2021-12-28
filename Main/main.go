package main

import (
	"fmt"

	Arrays "github.com/AndreoBouzas/Go_By_Example/Arrays"
	Constants "github.com/AndreoBouzas/Go_By_Example/Constants"
	For "github.com/AndreoBouzas/Go_By_Example/For"
	Hello_World "github.com/AndreoBouzas/Go_By_Example/Hello_World"
	If_Else "github.com/AndreoBouzas/Go_By_Example/If_Else"
	Slices "github.com/AndreoBouzas/Go_By_Example/Slices"
	Switch "github.com/AndreoBouzas/Go_By_Example/Switch"
	Values "github.com/AndreoBouzas/Go_By_Example/Values"
	Variables "github.com/AndreoBouzas/Go_By_Example/Variables"
	color "github.com/fatih/color"
)

var (
	underRed      = color.New(color.FgRed).Add(color.Underline)
	exampleNumber = 0
)

func main() {
	color.HiWhite("Digite o NÚMERO do exemplo que gostaria de ohar! :")
	underRed.Println("Exemplo 1  (Hello World!)")
	underRed.Println("Exemplo 2  (Valores e seus tipos)")
	underRed.Println("Exemplo 3  (Variáveis e Declarações)")
	underRed.Println("Exemplo 4  (Constantes)")
	underRed.Println("Exemplo 5  (Laço For)")
	underRed.Println("Exemplo 6  (If/Else)")
	underRed.Println("Exemplo 7  (Switch)")
	underRed.Println("Exemplo 8  (Arrays)")
	underRed.Println("Exemplo 9  (Slices)")
	fmt.Scan(&exampleNumber)

	switch exampleNumber {
	case 1:
		Hello_World.HelloWorld()
	case 2:
		Values.Values()
	case 3:
		Variables.Variables()
	case 4:
		Constants.Constants()
	case 5:
		For.For()
	case 6:
		If_Else.If_Else()
	case 7:
		Switch.Switch()
	case 8:
		Arrays.Arrays()
	case 9:
		Slices.Slices()
	}

}
