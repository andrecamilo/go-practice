// Go suporta [_funções anônimas_](http://en.wikipedia.org/wiki/Anonymous_function),
// que podem formar <a href="http://en.wikipedia.org/wiki/Closure_(computer_science)"><em>closures</em></a>.
// Funções anônimas são úteis quando você deseja definir
// uma função em linha sem ter que nomeá-lo.

package main

import "fmt"

// Essa função `intSeq` retorna outra função, que
// definimos anonimamente no corpo do `intSeq`. A
// função retornada _se fecha sobre_ a variável `i` para
// formar o closure.
func intSeq() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

func main() {

    // Chamamos `intSeq`, atribuindo o resultado (a função)
    // para o `nextInt`. Esse valor de função captura seu
    // próprio valor `i`, que será atualizado cada vez
    // que chamamos o `nextInt`.    
    nextInt := intSeq()

    // Veja o efeito do closure chamando o `nextInt`
    // algumas vezes.    
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    // Para confirmar que o estado é único para essa
    // função particular, crie e teste um novo.    
    newInts := intSeq()
    fmt.Println(newInts())
}