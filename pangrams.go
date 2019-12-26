package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func pangrams(s string) string {

	alfa := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	strs := []string{}
	s1 := strings.ReplaceAll(s, " ", "")
	s2 := strings.ToLower(s1)

	// transformar string em array
	for i := 0; i < len((s2)); i++ {
		strs = append(strs, string(s2[i]))
	}

	// repassar todas as letras do alfabeto
	for i := 0; i < len((alfa)); i++ {
		letra := alfa[i]

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
func PangramsInit() {
	//os.Stdin.WriteString("asasas")
	//reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	//s := readLine(reader)
	//We promptly judged antique ivory buckles for the next prize
	//We promptly judged antique ivory buckles for the prize
	//result := pangrams(s)

	result := pangrams("We promptly judged antique ivory buckles for the prize")

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
