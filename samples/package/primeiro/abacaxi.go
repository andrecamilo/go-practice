package primeiro

import (
	"fmt"

	r "github.com/crgimenes/rotateString"
)

//AbacaxiService é um serviço
type AbacaxiService struct{}

//CriarAbacaxi ex metodo 1
func (AbacaxiService) CriarAbacaxi() int {
	s := r.Rotate("Isto é um teste de Rotate")
	fmt.Printf("rotate: %s\r\n", s)

	return 2
}
