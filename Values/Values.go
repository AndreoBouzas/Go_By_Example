package Value

import (
	"fmt"

	Explain "github.com/AndreoBouzas/Go_By_Example/Explain"
)

func Values() {
	fmt.Println(Explain.ExplainThis(2))
	fmt.Print(` 
	
	func main() {

		Strings, Podem ser adicionadas com o +.
			fmt.Println("go" + "lang")


		Integers e floats.
			fmt.Println("1+1 =", 1+1)
			fmt.Println("7.0/3.0 =", 7.0/3.0)


		Booleans, uso dos operadores booleanos como esperado, true/ false
			fmt.Println(true && false)
			fmt.Println(true || false)
			fmt.Println(!true)
		}  
		` + "\n\n")

	fmt.Print("Os resultados são: \n\n")
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

}
