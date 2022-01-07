package main

import (
	"fmt"

	Arrays "github.com/AndreoBouzas/Go_By_Example/Arrays"
	ChannelBuffering "github.com/AndreoBouzas/Go_By_Example/ChannelBuffering"
	ChannelDirections "github.com/AndreoBouzas/Go_By_Example/ChannelDirections"
	ChannelSynchronization "github.com/AndreoBouzas/Go_By_Example/ChannelSynchronization"
	Channels "github.com/AndreoBouzas/Go_By_Example/Channels"
	Closures "github.com/AndreoBouzas/Go_By_Example/Closures"
	Constants "github.com/AndreoBouzas/Go_By_Example/Constants"
	Embedding "github.com/AndreoBouzas/Go_By_Example/Embedding"
	Errors "github.com/AndreoBouzas/Go_By_Example/Errors"
	For "github.com/AndreoBouzas/Go_By_Example/For"
	Functions "github.com/AndreoBouzas/Go_By_Example/Functions"
	Goroutines "github.com/AndreoBouzas/Go_By_Example/Goroutines"
	Hello_World "github.com/AndreoBouzas/Go_By_Example/Hello_World"
	If_Else "github.com/AndreoBouzas/Go_By_Example/If_Else"
	Interfaces "github.com/AndreoBouzas/Go_By_Example/Interfaces"
	Maps "github.com/AndreoBouzas/Go_By_Example/Maps"
	Methods "github.com/AndreoBouzas/Go_By_Example/Methods"
	Pointers "github.com/AndreoBouzas/Go_By_Example/Pointers"
	Range "github.com/AndreoBouzas/Go_By_Example/Range"
	Recursion "github.com/AndreoBouzas/Go_By_Example/Recursion"
	Select "github.com/AndreoBouzas/Go_By_Example/Select"
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
	underRed            = color.New(color.FgHiRed).Add(color.Underline)
)

func main() {

	color.HiWhite("\nDIGITE o NÚMERO do exemplo que gostaria de ohar! :\n")
	formatVibrantYellow.Println(`
	Exemplo 1   (Hello World!)		Exemplo 14  (Closures)
	Exemplo 2   (Valores e seus tipos)	Exemplo 15  (Recursion)	
	Exemplo 3   (Variáveis e Declarações)	Exemplo 16  (Ponteiros)
	Exemplo 4   (Constantes)		Exemplo 17  (Estruturas)
	Exemplo 5   (Laço For)			Exemplo 18  (Métodos)
	Exemplo 6   (If/Else)			Exemplo 19  (Interfaces)
	Exemplo 7   (Switch)			Exemplo 20  (Incorporação)
	Exemplo 8   (Arrays)			Exemplo 21  (Erros)
	Exemplo 9   (Slices)			Exemplo 22  (Goroutines)
	Exemplo 10  (Maps)			Exemplo 23  (Canais)
	Exemplo 11  (Range)			Exemplo 24  (Buffer de canal)
	Exemplo 12  (Functions)			Exemplo 25  (Sincronização de canais)
	Exemplo 13  (Função Variável)		Exemplo 26  (Canais Direcionáis)
						Exemplo 27  (Select)`)

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
	case 21:
		Errors.Errors()
	case 22:
		Goroutines.Goroutines()
	case 23:
		Channels.Channels()
	case 24:
		ChannelBuffering.ChannelBuffering()
	case 25:
		ChannelSynchronization.ChannelSynchronization()
	case 26:
		ChannelDirections.ChannelDirections()
	case 27:
		Select.Select()
	default:
		underRed.Print("\n\n\n!ATENÇÃO!\n\n\n!NÚMERO DE EXEMPLO INVÁLIDO!\n\n\n")

	}

}
