package HelloWorld //Declarando o nome do package

import ( //Importando os packages
	"fmt"

	Explain "github.com/AndreoBouzas/Go_By_Example/Explain"
)

//Declarando a função principal
func HelloWorld() {
	//chamando a função de explicação
	Explained := Explain.ExplainThis(1)
	//Print visual da explicação
	fmt.Println(Explained)
	//Print na tela da mensagem definida!
	fmt.Println("Hello World")
}
