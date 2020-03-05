package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/ahmetb/go-linq"
)

var index int
var setHeader bool

func main() {
	text := "/media/andrecamilo/Dados/dbs/locat/bases_locat/LOCAT_20200131.json"
	f, err := os.Open(text)
	reg, _ := regexp.Compile("[^a-zA-Z0-9-:{}._,\\[\"/\\]]+")
	// reader := bufio.NewReader(os.Stdin)
	// text, _ := reader.ReadString('\n')
	// text = strings.TrimSpace(text)
	// text = text + ".JSON"
	// println(text)
	// f, err := os.Open("./" + text)
	if err != nil {
		fmt.Println(err)
		return
	}
	stats, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	var size int64 = stats.Size()
	fmt.Println(size)
	data := make([]byte, 5000000)
	newFile, _ := os.Create(text + ".txt")
	defer newFile.Close()
	sobraJSON := ""
	for {
		data = data[:cap(data)]
		n, err := f.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}
		data = data[:n]
		processedString := reg.ReplaceAllString(string(data), "")
		str := strings.Split(processedString, "[")
		if len(str) > 1 {
			sobraJSON = readLines(str[1], sobraJSON, newFile)
		} else {
			str = strings.Split(processedString, "]")
			sobraJSON = readLines(str[0], sobraJSON, newFile)
		}
	}
	keys, _ := getKeys(sobraJSON)
	line, _ := formatLine(sobraJSON, keys)
	newFile.WriteString(line)
	fmt.Println("--- TERMINADO ---")
}
func readLines(str, sobra string, file *os.File) string {
	ch := make(chan string)
	defer close(ch)
	if len(str) > 0 && string(str[:2]) == ",{" {
		str = str[1:]
	}
	lsobra := sobra + str
	strRows := strings.ReplaceAll(lsobra, "},", "}@")
	rows := strings.Split(strRows, "@")
	countRows := len(rows) - 1
	keys, err := getKeys(rows[0])
	if err != nil {
		return str
	}
	if !setHeader {
		header := ""
		for _, item := range keys {
			header += fmt.Sprintf("%v;", item)
		}
		file.WriteString(header + "\n")
		setHeader = true
	}
	for i, r := range rows {
		go func(i int, r string) {
			if countRows == i {
				ch <- r
				return
			}
			line, _ := formatLine(r, keys)
			_, err := file.WriteString(line)
			if err != nil {
				fmt.Println(err)
			}
		}(i, r)
	}
	s := <-ch
	index += countRows + 1
	fmt.Println(fmt.Sprintf("Gravado linhas -> %v", index))
	file.Sync()
	return s
}
func formatLine(r string, keys []string) (string, error) {
	data := map[string]interface{}{}
	err := json.Unmarshal([]byte(r), &data)
	line := ""
	for _, item := range keys {
		if item == "data" {
			dt := data[item].(string)
			minutes := dt[len(dt)-5:]
			dmy := dt[:10]
			line += fmt.Sprintf("%v %v;", dmy, minutes)
		} else {
			line += fmt.Sprintf("%v;", data[item])
		}
	}
	return line + "\n", err
}
func getKeys(line string) ([]string, error) {
	data := map[string]interface{}{}
	err := json.Unmarshal([]byte(line), &data)
	if err != nil {
		return nil, err
	}
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	linq.From(keys).OrderByT(func(s string) string { return s }).ToSlice(&keys)
	return keys, nil
}
