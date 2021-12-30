package main

import (
	"fmt"

	Arrays "github.com/AndreoBouzas/Go_By_Example/Arrays"
	Constants "github.com/AndreoBouzas/Go_By_Example/Constants"
	For "github.com/AndreoBouzas/Go_By_Example/For"
	Hello_World "github.com/AndreoBouzas/Go_By_Example/Hello_World"
	If_Else "github.com/AndreoBouzas/Go_By_Example/If_Else"
	Maps "github.com/AndreoBouzas/Go_By_Example/Maps"
	Slices "github.com/AndreoBouzas/Go_By_Example/Slices"
	Switch "github.com/AndreoBouzas/Go_By_Example/Switch"
	Values "github.com/AndreoBouzas/Go_By_Example/Values"
	Variables "github.com/AndreoBouzas/Go_By_Example/Variables"
	color "github.com/fatih/color"
)

var (
	formatVibrantYellow = color.New(color.FgHiYellow)
	formatVibrantBlack  = color.New(color.FgBlack)
	exampleNumber       = 0
)

func main() {

	color.HiWhite("Digite o NÚMERO do exemplo que gostaria de ohar! :")
	formatVibrantYellow.Println("Exemplo 1   (Hello World!)")
	formatVibrantYellow.Println("Exemplo 2   (Valores e seus tipos)")
	formatVibrantYellow.Println("Exemplo 3   (Variáveis e Declarações)")
	formatVibrantYellow.Println("Exemplo 4   (Constantes)")
	formatVibrantYellow.Println("Exemplo 5   (Laço For)")
	formatVibrantYellow.Println("Exemplo 6   (If/Else)")
	formatVibrantYellow.Println("Exemplo 7   (Switch)")
	formatVibrantYellow.Println("Exemplo 8   (Arrays)")
	formatVibrantYellow.Println("Exemplo 9   (Slices)")
	formatVibrantYellow.Println("Exemplo 10  (Maps)")
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
	case 10:
		Maps.Maps()
	}

}
