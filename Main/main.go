package main

import (
	"fmt"

	Hello_World "github.com/AndreoBouzas/Go_By_Example/Hello_World"
	Values "github.com/AndreoBouzas/Go_By_Example/Values"
)

func main() {
	fmt.Println("Digite o número do exemplo que gostaria de ohar! :")
	exampleNumber := 0
	fmt.Scan(&exampleNumber)

	if exampleNumber == 1 {
		Hello_World.HelloWorld()
	} else if exampleNumber == 2 {
		Values.Values()
	}

}
