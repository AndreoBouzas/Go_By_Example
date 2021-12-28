package HelloWorld //Declarando o nome do package

import ( //Importando os packages

	Explain "github.com/AndreoBouzas/Go_By_Example/Explain"
	color "github.com/fatih/color"
)

var (
	underBlue    = color.New(color.FgCyan)
	underRed     = color.New(color.FgHiRed).Add(color.Underline)
	exampleStyle = color.New(color.FgHiWhite)
)

//Declarando a função principal
func HelloWorld() {
	underRed.Print("Exemplo 1  (Hello World!): \n\n")
	//chamando a função de explicação
	underBlue.Println(Explain.ExplainThis(1))
	//Print na tela da mensagem definida!
	exampleStyle.Println("Hello World")
}
