package HelloWorld //Declarando o nome do package

import ( //Importando os packages
	"fmt"

	Explain "github.com/AndreoBouzas/Go_By_Example/Explain"
)

//Declarando a função principal
func HelloWorld() {
	fmt.Print("Exemplo 1  (Hello World!): \n\n")
	//chamando a função de explicação
	fmt.Println(Explain.ExplainThis(1))
	//Print na tela da mensagem definida!
	fmt.Println("Hello World")
}
