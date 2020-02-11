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

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {

	var aa int32
	var bb int32

	for i := 0; i < len(arr[1]); i++ {
		v := arr[i][i]
		aa += int32(v)
	}

	l := len(arr[1]) - 1
	for i := l; i >= 0; i-- {
		v := arr[i][i]
		bb += int32(v)
	}

	f := float64(aa - bb)
	return int32(math.Abs(f))

}

func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	//nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	//n := int32(nTemp)
	n := 3

	var arr [][]int32
	/*
		for i := 0; i < int(n); i++ {
			arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var arrRow []int32
			for _, arrRowItem := range arrRowTemp {
				arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
				checkError(err)
				arrItem := int32(arrItemTemp)
				arrRow = append(arrRow, arrItem)
			}

			if len(arrRow) != int(n) {
				panic("Bad input")
			}

			arr = append(arr, arrRow)
		}*/

	for i := 0; i < int(n); i++ {
		var arrRowTemp string

		if i == 0 {
			arrRowTemp := strings.Split(strings.TrimRight("11 2 4", " \t\r\n"), " ")
		}
		if i == 1 {
			arrRowTemp := strings.Split(strings.TrimRight("4 5 6", " \t\r\n"), " ")
		}
		if i == 2 {
			arrRowTemp := strings.Split(strings.TrimRight("10 8 -12", " \t\r\n"), " ")
		}

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(n) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := diagonalDifference(arr)

	fmt.Fprintf(writer, "%d\n", result)

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
