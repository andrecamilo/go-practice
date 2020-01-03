// Nós podemos usar canais para sincronizar a execução
// através das goroutines. Aqui temos um exemplo do uso
// de um bloqueador de recebimento para esperar a finalização
// de uma goroutine.

package main

import "fmt"
import "time"

// Essa é a função que nós iremos rodar em uma goroutine. O
// canal `feito` será usado para notificar outra
// goroutine que essa função trabalho é feita.
func trabalhador(feito chan bool) {
	fmt.Print("trabalhando...")
	time.Sleep(time.Second)
	fmt.Println("feito")

	// Envia um valor para notificar o que estamos fazendo.
	feito <- true
}

func main() {

	// Inicia uma goroutine trabalhador, dando-lhe o canal para
	// notificar.
	feito := make(chan bool, 1)
	go trabalhador(feito)

	// Bloqueia até que recebamos uma notificação do
	// trabalhador no canal.
	<-feito
}
