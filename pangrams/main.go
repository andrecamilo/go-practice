package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func pangrams(s string) string {

	alfabeto := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	strs := []string{}

	tratato :=
		strings.ToLower(
			strings.ReplaceAll(s, " ", ""))

	// transformar string em array
	for i := 0; i < len((tratato)); i++ {
		strs = append(strs, string(tratato[i]))
	}

	// repassar todas as letras do alfabeto
	for i := 0; i < len((alfabeto)); i++ {
		letra := alfabeto[i]

		//procurar letra do alfabeto na frase
		found := false
		for key, value := range strs {
			if value == letra {
				found = true
				fmt.Printf(string(key))
			}
		}

		if found == false {
			return "not pangram"
		}
	}

	return "pangram"
}

// PangramsInit inicia um algoritmo que verifica se
// uma frase contem todas as letras do alfabeto
func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	if err != nil {

	}
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	fmt.Println("digite uma frase")
	//s := readLine(reader)

	//result := pangrams("We promptly judged antique ivory buckles for the next prize")
	result := pangrams("We promptly judged antique ivory buckles for the prize")
	//result := pangrams(s)

	fmt.Print(result)
	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {

	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
