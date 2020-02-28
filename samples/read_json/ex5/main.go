package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

func readLines(filename string) ([]string, error) {
	var lines []string
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return lines, err
	}
	buf := bytes.NewBuffer(file)
	for {
		line, err := buf.ReadString('\n')
		if len(line) == 0 {
			if err != nil {
				if err == io.EOF {
					break
				}
				return lines, err
			}
		}
		lines = append(lines, line)
		if err != nil && err != io.EOF {
			return lines, err
		}
	}
	return lines, nil
}

func main() {
	// a flat file that has 339276 lines of text in it for a size of 62.1 MB

	//fileName := "../teste.json"
	//fileName := "/media/andrecamilo/Dados/dbs/locacao/bases_locacao/locacao_20200131.json"
	fileName := "/media/andrecamilo/Dados/dbs/locacao/bases_locacao/locacao_JAN_01_10.JSON"
	lines, err := readLines(fileName)
	fmt.Println(len(lines))
	if err != nil {
		fmt.Println(err)
		return
	}
}
