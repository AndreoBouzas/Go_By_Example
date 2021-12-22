package Constants

import (
	"fmt"
	"math"

	Explain "github.com/AndreoBouzas/Go_By_Example/Explain"
)

const s string = "constante"

func Constants() {
	fmt.Print("Exemplo 4  (Constantes): \n\n")
	fmt.Println(Explain.ExplainThis(4))
	fmt.Print("const declara um valor constante. \n " + " Uma declaração const  pode aparecer em qualquer lugar que uma declaração var  possa. " + ":EXEMPLO : \n" + `
	const  s  string  =  "constante"
função  main ()  { 
    fmt.Println(s)
	const  n  =  500000000` + "\n\n RESULTADO : \n\n")
	const n = 500000000
	fmt.Print(s, "\n\n")

	fmt.Print("Expressões constantes realizam operações aritméticas com precisão arbitrária. : \n EXEMPLO : \n" + `
	const  d  =  3e20  /  n 
    fmt.Println(d)` + "\n\n RESULTADO : \n\n")
	const d = 3e20 / n
	fmt.Print(d, "\n\n")

	fmt.Print("Uma constante numérica não tem tipo até que seja fornecida, como por meio de uma conversão explícita. :\n EXEMPLO : \n" + `
    fmt.Println(int64(d))` + "\n\n RESULTADO : \n\n")
	fmt.Print(int64(d), "\n\n")

	fmt.Print("Um número pode receber um tipo usando-o em um contexto que requer um, como uma atribuição de variável ou chamada de função. Por exemplo, aqui math.Sin espera um float64.: \n EXEMPLO : \n" + `
	fmt.Println(math.Sin(n)) ` + "\n\n RESULTADO : \n\n")
	fmt.Print(math.Sin(n), "\n\n")
}
