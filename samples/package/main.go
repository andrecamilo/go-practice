package main

import (
	"fmt"
	"strconv"

	"package/primeiro"
	"package/primeiro/segundo"
)

func main() {
	fmt.Printf("teste")

	a := primeiro.AbacateService{}

	x := a.CriarAbacate()
	y := segundo.CriarLaranja()

	fmt.Printf(strconv.Itoa(y))
	fmt.Printf(strconv.Itoa(x))
}
