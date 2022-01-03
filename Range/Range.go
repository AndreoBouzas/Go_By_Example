package Range

import (
	"fmt"

	Explain "github.com/AndreoBouzas/Go_By_Example/Explain"
	color "github.com/fatih/color"
)

var (
	underBlue    = color.New(color.FgCyan)
	underRed     = color.New(color.FgHiRed).Add(color.Underline)
	exampleStyle = color.New(color.FgHiWhite)
	exampleText  = color.New(color.FgHiGreen)
	resultText   = color.New(color.FgHiMagenta)
)

func Range() {
	underRed.Print("\n\n\n Exemplo 11  (Range) \n\n")
	underBlue.Printf(Explain.ExplainThis(11))
	fmt.Print("\n Aqui, usamos range para somar os números em um Slice. Arrays funcionam assim também. \n")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	`)
	resultText.Print("\nRESULTADO : \n")
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//----------------------------------------------------------------------
	underBlue.Print("\n Range sobre Arrays e Slices fornece o índice e o valor de cada entrada. Acima não precisamos do índice, então o ignoramos com o identificador em branco _. Às vezes, queremos realmente os índices.   \n")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	`)
	resultText.Print("\nRESULTADO : \n")
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//----------------------------------------------------------------------
	underBlue.Print("\n Range sobre Maps itera sobre pares de chave valor\n")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%S -> %S\n", k, v)
	}
	`)
	resultText.Print("\nRESULTADO : \n")
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//----------------------------------------------------------------------
	underBlue.Print("\n Range também pode iterar sobre apenas as chaves de um Map\n")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	for k := range kvs {
		fmt.Println("key:", k)
	}
	`)
	resultText.Print("\nRESULTADO : \n")
	for k := range kvs {
		fmt.Println("key:", k)
	}
	//----------------------------------------------------------------------
	underBlue.Print("\n Range em strings itera sobre pontos de código Unicode. O primeiro valor é o índice de bytes inicial do rune e o segundo o rune próprio.\n")
	exampleText.Print("\nEXEMPLO : \n")
	exampleStyle.Print(`
	for i, c := range "go" {
		fmt.Println(i, c)
	}
	`)
	resultText.Print("\nRESULTADO : \n")
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
