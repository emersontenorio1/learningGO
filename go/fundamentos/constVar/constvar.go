package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2
	var nome = "emerson"
	// forma reduzida de criar uma var

	area := PI * math.Pow(raio, 2)

	fmt.Println(nome, "o raio é", area)

	const (
		a     = 1
		b     = 2
		nome2 = "emerson"
	)

	var (
		c = 3
		d = 5
	)

	fmt.Println(a, b, c, d, nome2)

	var e, f = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "epa"

	fmt.Println(g, h, i)

}
