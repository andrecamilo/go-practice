package primeiro

import (
	"fmt"

	r "github.com/crgimenes/rotateString"
)

//TomateService é um serviço
type TomateService struct{}

//CriarTomate ex metodo 1
func (TomateService) CriarTomate() int {
	s := r.Rotate("Isto é um teste de Rotate")
	fmt.Printf("rotate: %s\r\n", s)

	return 2
}
