/*
func main() {
	ch := make(chan int, 1)
	ch <- 1 // enviando dados para o canal (escrita)
	<-ch    // recebendo dados do canal (leitura)
	ch <- 2
	fmt.Println(<-ch)
}
*/

package main

import (
	"fmt"
	"time"
)

// Channel (canal) - é a forma de comunicação entre goroutines
// channel é um tipo
func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal
	time.Sleep(time.Second)
	c <- 3 * base
	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)

	//ok
	// a, b := <-c, <-c // recebendo dados do canal
	// fmt.Println(a, b)

	//ok
	x, y, z := <-c, <-c, <-c // recebendo dados do canal
	fmt.Println(x, y, z)

	//ok
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	//err
	//fmt.Println(<-c)
}
