package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Funcao que le o conteudo do arquivo e retorna um slice the string com todas as linhas do arquivo
func lerTexto(caminhoDoArquivo string) ([]string, error) {
	// Abre o arquivo
	arquivo, err := os.Open(caminhoDoArquivo)
	// Caso tenha encontrado algum erro ao tentar abrir o arquivo retorne o erro encontrado
	if err != nil {
		return nil, err
	}
	// Garante que o arquivo sera fechado apos o uso
	defer arquivo.Close()

	// Cria um scanner que le cada linha do arquivo
	var linhas []string
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linhas = append(linhas, scanner.Text())
	}

	// Retorna as linhas lidas e um erro se ocorrer algum erro no scanner
	return linhas, scanner.Err()
}

// Funcao que escreve um texto no arquivo e retorna um erro caso tenha algum problema
func escreverTexto(linhas []string, caminhoDoArquivo string) error {
	// Cria o arquivo de texto
	arquivo, err := os.Create(caminhoDoArquivo)
	// Caso tenha encontrado algum erro retornar ele
	if err != nil {
		return err
	}
	// Garante que o arquivo sera fechado apos o uso
	defer arquivo.Close()

	// Cria um escritor responsavel por escrever cada linha do slice no arquivo de texto
	escritor := bufio.NewWriter(arquivo)
	for _, linha := range linhas {
		fmt.Fprintln(escritor, linha)
	}

	// Caso a funcao flush retorne um erro ele sera retornado aqui tambem
	return escritor.Flush()
}

func main() {
	var conteudo []string
	conteudo = append(conteudo, "123")
	conteudo = append(conteudo, "456")
	conteudo = append(conteudo, "789")

	err := escreverTexto(conteudo, "foo.txt")
	if err != nil {
		log.Fatalf("Erro:", err)
	}

	conteudo = nil
	conteudo, err = lerTexto("foo.txt")
	if err != nil {
		log.Fatalf("Erro:", err)
	}

	for indice, linha := range conteudo {
		fmt.Println(indice, linha)
	}
}
