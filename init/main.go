package main

import "fmt"

func init() {
	fmt.Println("Inicializando...")

	i := 1
	var p *int = nil
	p = &i // pegando o endereço da variável
	*p++   // desreferenciando (pegando o valor)
	i++
	// Go não tem aritmética de ponteiros
	// p++
	fmt.Println(p, *p, i, &i)

}

func main() {
	fmt.Println("Main...")

	// 	//NestedInit()
	// 	//Callbyreference()
	// 	//Callbyreference2()
	// 	//Slicesort()
	// 	PangramsInit()
}
