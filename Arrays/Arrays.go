package Arrays

import (
	"fmt"

	Explain "github.com/AndreoBouzas/Go_By_Example/Explain"
)

func Arrays() {
	fmt.Print("Exemplo 8  (Arrays): \n\n")
	fmt.Println(Explain.ExplainThis(8))

	fmt.Print("criamos um array aque conterá exatamente 5 ints. O tipo de elementos e comprimento são parte do tipo da matriz. Por padrão, uma matriz tem valor zero, o que para ints significa 0s. \n EXEMPLO : \n" + `
	
	var a [5]int
    fmt.Println("emp:", a)
	` + "\n\n RESULTADO : \n\n")
	var a [5]int
	fmt.Print("emp:", a, "\n\n")

	fmt.Print("Podemos definir um valor em um índice usando a array[índice] = Valor sintaxe e obter um valor com array[índice]. \n EXEMPLO : \n" + `
	
	a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])
	` + "\n\n RESULTADO : \n\n")
	a[4] = 100
	fmt.Print("set:", a, "\n\n")
	fmt.Print("get:", a[4], "\n\n")

	fmt.Print("O builtin len retorna o comprimento de uma matriz. \n EXEMPLO : \n" + `
	fmt.Println("len:", len(a))
	` + "\n\n RESULTADO : \n\n")
	fmt.Println("len:", len(a))

	fmt.Print("podemos usar a seguinte sintaxe para declarar e inicializar um array em uma linha. \n EXEMPLO : \n" + `
	b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)
	` + "\n\n RESULTADO : \n\n")
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	fmt.Print("Os tipos de matriz são unidimensionais, mas você pode compor tipos para construir estruturas de dados multidimensionais. \n EXEMPLO : \n" + `
	var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
	` + "\n\n RESULTADO : \n\n")

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
