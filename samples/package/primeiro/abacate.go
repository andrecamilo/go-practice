package primeiro

import (
	"fmt"
	"package/primeiro/segundo"
	"strconv"

	r "github.com/crgimenes/rotateString"
)

//AbacateService é um servico
type AbacateService struct{}

//CriarAbacate ex metodo 1
func (AbacateService) CriarAbacate() int {
	s := r.Rotate("Isto é um teste de Rotate")
	fmt.Printf("rotate: %s\r\n", s)

	//a := AbacaxiService{}
	//x := a.CriarAbacaxi()

	l := segundo.CriarLaranja()

	fmt.Printf(strconv.Itoa(l))
	return 2
}
