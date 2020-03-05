package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//fileName := "/media/andrecamilo/Dados/dbs/locacao/bases_locacao/locacao_20200131.json"
	fileName := "teste.json"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
