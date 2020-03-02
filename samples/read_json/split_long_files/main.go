package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//f, err := os.Open("locacao_20200131_bkp.json")
	f, err := os.Open("/media/andrecamilo/Dados/dbs/locacao/bases_locacao/locacao_JAN_01_10_1.JSON")
	//longstring := "This is a very long string. Not."
	handle(err)

	var words []string
	//scanner := bufio.NewScanner(strings.NewReader(longstring))
	scanner := bufio.NewScanner(bufio.NewReader(f))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	fmt.Println("word list:")
	for _, word := range words {
		fmt.Println(word)
	}
}

func handle(err error) {
	//panic(err)
}
