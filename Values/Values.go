package Value

import (
	"fmt"

	Explain "github.com/AndreoBouzas/Go_By_Example/Explain"
)

func Values() {
	fmt.Print("Exemplo 2  (Valores e seus tipos): \n\n")
	fmt.Println(Explain.ExplainThis(2))
	fmt.Print("Strings, Podem ser adicionadas com o + \n EXEMPLO : \n" + `
	fmt.Println("go" + "lang")` + "\n\n RESULTADO : \n\n")
	fmt.Print("go" + "lang\n\n")
	fmt.Print("Integers e floats.\n EXEMPLO : \n" + `
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)` + "\n\n RESULTADO : \n\n")
	fmt.Println("1+1 =", 1+1)
	fmt.Print("7.0/3.0 =", 7.0/3.0, "\n\n")
	fmt.Print("Booleans, uso dos operadores booleanos como esperado, true/ false \n EXEMPLO : \n" + `
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)` + "\n\n RESULTADO : \n\n")
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
