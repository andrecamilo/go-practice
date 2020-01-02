package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the funnyString function below.
func funnyString(s string) string {

	funny := true
	runes := []rune(s)
	j := len(runes) - 1

	for i := 1; i < len(s); i++ {

		xx := math.Abs(float64(runes[i]) - float64(runes[i-1]))
		yy := math.Abs(float64(runes[j]) - float64(runes[j-1]))

		if xx != yy {
			funny = false
			break
		}

		j--
	}

	if funny == true {
		return "Funny"
	}

	return "Not Funny"
}

/*
func funnyString(s string) string {
	fmt.Println("iniciando funnyString")

	runes := []rune(s)
	var result []int
	var diferencas []int
	var reverseString string

	for i := 0; i < len(runes); i++ {
		reverseString = string(runes[i]) + reverseString
	}

	for i := 0; i < len(runes); i++ {
		result = append(result, int(runes[i]))
	}

	resultReverse := make([]int, len(result))
	for i, j := 0, len(result)-1; i <= j; i, j = i+1, j-1 {
		resultReverse[i], resultReverse[j] = result[j], result[i]
	}

	for i := 0; i < len(result); i++ {
		if (i + 1) < len(result) {
			x := result[i] - result[i+1]
			diferencas = append(diferencas, x)
		}
	}

	diferencasReverse := make([]int, len(diferencas))
	for i, j := 0, len(diferencas)-1; i <= j; i, j = i+1, j-1 {
		diferencasReverse[i], diferencasReverse[j] = diferencas[j], diferencas[i]
	}

	fmt.Println("String: ")
	fmt.Println(s)

	fmt.Println("String reverso: ")
	fmt.Println(reverseString)

	fmt.Println("ASCII: ")
	fmt.Println(result)

	fmt.Println("ASCII reverso: ")
	fmt.Println(resultReverse)

	fmt.Println("Diferencas ")
	fmt.Println(diferencas)

	fmt.Println("Diferencas reverso: ")
	fmt.Println(diferencasReverse)

	if reflect.DeepEqual(diferencasReverse, diferencas) {
		return "Funny"
	}

	return "Not Funny"
}
*/

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := funnyString(s)

		fmt.Fprintf(writer, "%s\n", result)
		fmt.Println(result)
	}

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
