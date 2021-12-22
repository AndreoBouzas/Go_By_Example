package Variables

import (
	"fmt"

	Explain "github.com/AndreoBouzas/Go_By_Example/Explain"
)

func Variables() {
	fmt.Print("Exemplo 3  (Variáveis e Declarações): \n\n")
	fmt.Println(Explain.ExplainThis(3))
	fmt.Print("Var declara 1 ou mais variáveis \n EXEMPLO : \n" + `
	var a = "Andreo"
	fmt.Println(a)` + "\n\n RESULTADO : \n\n")
	var a = "Andreo"
	fmt.Print(a, "\n\n")

	fmt.Print("Declarando multiplas vaiáveis com uso do var: \n EXEMPLO : \n" + `
	var b, c int = 1, 2
	fmt.Println(b, c)` + "\n\n RESULTADO : \n\n")
	var b, c int = 1, 2
	fmt.Print(b, c, "\n\n")

	fmt.Print("Go associa um tipo a uma variável, basedo no valor atribuído a ela. Caso o tipo da mesma não tenha sido declarado! \n\n")
	fmt.Print("Declaração de variável do tipo booleano :\n EXEMPLO : \n" + `
	var d = true
    fmt.Println(d)` + "\n\n RESULTADO : \n\n")
	var d = true
	fmt.Print(d, "\n\n")

	fmt.Print("Variáveis declaradas sem valor de inicialização, recebem por padrão o zero do tipo em questão: \n EXEMPLO : \n" + `
	var e int
    fmt.Println(e)` + "\n\n RESULTADO : \n\n")
	var e int
	fmt.Print(e, "\n\n")

	fmt.Print("O atalho de inicialização e declaração de variáveis é :=    \n EXEMPLO : \n" + `
	f := "Maçã"
    fmt.Println(f)` + "\n\n RESULTADO : \n\n")
	f := "Maçã"
	fmt.Print(f, "\n\n")

}
