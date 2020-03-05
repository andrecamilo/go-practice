// _Funções_ são fundamentais no Go. Vamos aprender sobre
// funções com alguns exemplos diferentes.

package main

import "fmt"

// Aqui temos uma função que pega dois `int`eiros e retorna
// sua soma como um `int`.
func mais(a int, b int) int {
     
    // Go exige retornos explícitos, ou seja, ele não vai
    // retornar automaticamente o valor da última
    // expressão.    
    return a + b
}

func main() {

    // Chama uma função como esperado, com
    // `nome(args)`.    
    res := mais(1, 2)
    fmt.Println("1+2 =", res)
}

// todo: tipos de parâmetros aglutinados