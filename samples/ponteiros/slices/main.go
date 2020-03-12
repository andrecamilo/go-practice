// _Slices_ são tipos de dados fundamentais no Go, fornecendo uma
// interface mais poderosa para seqüências de matrizes.

package main

import "fmt"

func main() {

    // Ao contrário das matrizes, os slices são digitados apenas
    // pelos elementos que eles contêm (e não o número de elementos).
    // Para criar um slice vazio com tamanho diferente de zero, use
    // o `make` padrão. Aqui nós fizemos um slice de uma `string` de 
    // tamanho `3` (inicialmente com valor zero).
    
    s := make([]string, 3)
    fmt.Println("emp:", s)

    // Podemos definir e obter assim como as matrizes.    
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    // O `len` retorna o tamanho do slice, como esperado.    
    fmt.Println("len:", len(s))

    // Além das operações básicas, os slices suportam
    // muito mais, o que os tornam mais ricos do que as
    // matrizes. Uma é o padrão `append`, na qual
    // retorna um slice contendo um ou mais novos valores.
    // Note que precisamos aceitar um valor de retorno do
    // append assim como podemos obter um novo valor do slice.
    
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    // Slices podem ser também copiados. Aqui nós criamos um
    // slice vazio `c` do mesmo tamanho de `s` e copiamos
    // o `s` no `c`.    
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    // Slices suportam um operador "slice" com a sintaxe
    // `slice[min:max]`. Por exemplo, este recebe um slice
    // de elementos `s[2]`, `s[3]`, e `s[4]`.    
    l := s[2:5]
    fmt.Println("sl1:", l)

    // Esses slices vão até (mas excluindo) `s[5]`.    
    l = s[:5]
    fmt.Println("sl2:", l)

    // E esses slices vão do (e incluindo) `s[2]`.    
    l = s[2:]
    fmt.Println("sl3:", l)

    // Nós também podemos declarar e inicializar a variável para slice
    // em uma simples linha.    
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    // Slices podem ser compostas em estruturas de dados 
    // multi-dimensionais. O tamanho dos slices interiores pode
    // variar, ao contrário das matrizes multi-dimensionais.
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
