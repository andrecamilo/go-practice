package main

import (
	"fmt"
	"github.com/bcicen/jstream"
	"os"
)

func main() {
	//fileName := "/media/andrecamilo/Dados/dbs/locacao/bases_locacao/locacao_JAN_01_10.JSON"
	//fileName := "/home/andrecamilo/Projects/locacao_20200131.json"
	fileName := "../teste.json"
	f, _ := os.Open(fileName)
	decoder := jstream.NewDecoder(f, 1) // extract JSON values at a depth level of 1
	for mv := range decoder.Stream() {
		fmt.Printf("%v\n ", mv.Value)
	}
}
