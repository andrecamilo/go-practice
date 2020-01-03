// _range_ itera com elementos em uma variedade de
// estruturas de dados. Vamos ver como usar o `range` com
// algumas das estruturas de dados que já aprendemos.

package main

import "fmt"

func main() {

    // Aqui nós usamos o `range` para somar os números em um slice.
    // Matrizes trabalham da mesma maneira também.    
    nums := []int{2, 3, 4}
    soma := 0
    for _, num := range nums {
        soma += num
    }
    fmt.Println("soma:", soma)

    // `range` em matrizes ou slices fornece ambos o
    // index e o valor para cada entrada. Acima nós não
    // precisamos do index, então nós ignoramos ele com o 
    // _identificador em branco_ `_`. Às vezes, entretanto,
    // nós realmente queremos os indexes.    
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    // `range` no map itera com os pares chave/valor.    
    cvs := map[string]string{"a": "maçã", "b": "banana"}
    for c, v := range cvs {
        fmt.Printf("%s -> %s\n", c, v)
    }

    // `range` em strings itera com pontos do código Unicode.
    // O primeiro valor é o byte inicial, index do `rune`
    // e o segundo o próprio `rune`.    
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}