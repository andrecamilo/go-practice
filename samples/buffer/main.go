package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	chunksize int = 1024
)

var (
	data  *os.File
	part  []byte
	err   error
	count int
)

func openFile(name string) (byteCount int, buffer *bytes.Buffer) {

	data, err = os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	reader := bufio.NewReader(data)
	buffer = bytes.NewBuffer(make([]byte, 0))
	part = make([]byte, chunksize)

	for {
		if count, err = reader.Read(part); err != nil {
			break
		}
		buffer.Write(part[:count])
	}
	if err != io.EOF {
		log.Fatal("Error Reading ", err)
	} else {
		err = nil
	}

	byteCount = buffer.Len()
	return
}

func main() {
	fileName := "/media/andrecamilo/Dados/dbs/locacao/bases_locacao/locacao_20200131.json"
	length, _ := openFile(fileName)
	fmt.Println(length)
}
