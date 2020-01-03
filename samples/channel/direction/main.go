// Ao utilizar canais como parâmetros de funções, nós podemos
// especificar se o canal é destinado a apenas enviar ou receber
// valores. Essa especificidade aumenta a segurança de tipos
// do programa.

package main

import "fmt"

// Essa função `ping` aceita apenas um canal para enviar
// valores. Seria um erro de compilação-tempo tentar
// receber nesse canal.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// A função `pong` aceita um canal para receber
// (`pings`) e um segundo para enviar (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "mensagem passada")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
