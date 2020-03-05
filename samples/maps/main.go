// Maps são [tipos de dados associativos](http://en.wikipedia.org/wiki/Associative_array) padrões do Go
// (às vezes chamados _hashes_ ou _dicts_ em outras linguagens).

package main

import "fmt"

func main() {

    // Para criar um map vazio, use o padrão `make`:
    // `make(map[tipo-chave]tipo-val)`.    
    m := make(map[string]int)

    // Configura pares chave/valor usando a típica sintaxe `nome[chave] = val`.
    m["k1"] = 7
    m["k2"] = 13

    // Imprimir um map com, por ex. `Println` mostra todos
    // os pares chave/valor.   
    fmt.Println("map:", m)

    // Pega um valor para uma chave com o `nome[chave]`.    
    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    // O padrão `len` retorna o número do par chave/valor
    // quando chamado por um map.    
    fmt.Println("len:", len(m))

    // O padrão `delete` remove os pares chave/valor de
    // um map.    
    delete(m, "k2")
    fmt.Println("map:", m)

    // O segundo opcional, retorna um valor quando chega
    // o valor de um map, indicando se a chave estava presente no map.
    // Isso pode ser usado para distinguir entre as chaves perdidas
    // e as chaves com valor zero, como `0` ou `""`.    
    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    // Você pode também declarar ou iniciar um novo map na
    // mesma linha com essa sintaxe.    
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}