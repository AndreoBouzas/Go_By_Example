package main

import (
	"fmt"

	Arrays "github.com/AndreoBouzas/Go_By_Example/Arrays"
	Closures "github.com/AndreoBouzas/Go_By_Example/Closures"
	Constants "github.com/AndreoBouzas/Go_By_Example/Constants"
	Embedding "github.com/AndreoBouzas/Go_By_Example/Embedding"
	For "github.com/AndreoBouzas/Go_By_Example/For"
	Functions "github.com/AndreoBouzas/Go_By_Example/Functions"
	Hello_World "github.com/AndreoBouzas/Go_By_Example/Hello_World"
	If_Else "github.com/AndreoBouzas/Go_By_Example/If_Else"
	Interfaces "github.com/AndreoBouzas/Go_By_Example/Interfaces"
	Maps "github.com/AndreoBouzas/Go_By_Example/Maps"
	Methods "github.com/AndreoBouzas/Go_By_Example/Methods"
	Pointers "github.com/AndreoBouzas/Go_By_Example/Pointers"
	Range "github.com/AndreoBouzas/Go_By_Example/Range"
	Recursion "github.com/AndreoBouzas/Go_By_Example/Recursion"
	Slices "github.com/AndreoBouzas/Go_By_Example/Slices"
	Structs "github.com/AndreoBouzas/Go_By_Example/Structs"
	Switch "github.com/AndreoBouzas/Go_By_Example/Switch"
	Values "github.com/AndreoBouzas/Go_By_Example/Values"
	Variables "github.com/AndreoBouzas/Go_By_Example/Variables"
	VariadicFunctions "github.com/AndreoBouzas/Go_By_Example/VariadicFunctions"
	color "github.com/fatih/color"
)

var (
	formatVibrantYellow = color.New(color.FgHiYellow)
	formatVibrantBlack  = color.New(color.FgBlack)
	exampleNumber       = 0
)

func main() {

	color.HiWhite("\nDIGITE o NÚMERO do exemplo que gostaria de ohar! :\n")
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
	formatVibrantYellow.Println("Exemplo 11  (Range)")
	formatVibrantYellow.Println("Exemplo 12  (Functions)")
	formatVibrantYellow.Println("Exemplo 13  (Função Variável)")
	formatVibrantYellow.Println("Exemplo 14  (Closures)")
	formatVibrantYellow.Println("Exemplo 15  (Recursion)")
	formatVibrantYellow.Println("Exemplo 16  (Ponteiros)")
	formatVibrantYellow.Println("Exemplo 17  (Estruturas)")
	formatVibrantYellow.Println("Exemplo 18  (Métodos)")
	formatVibrantYellow.Println("Exemplo 19  (Interfaces)")
	formatVibrantYellow.Println("Exemplo 20  (incorporação)")
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
	case 11:
		Range.Range()
	case 12:
		Functions.Functions()
	case 13:
		VariadicFunctions.VariadicFunctions()
	case 14:
		Closures.Closures()
	case 15:
		Recursion.Recursion()
	case 16:
		Pointers.Pointers()
	case 17:
		Structs.Structs()
	case 18:
		Methods.Methods()
	case 19:
		Interfaces.Interfaces()
	case 20:
		Embedding.Embedding()
	}

}
