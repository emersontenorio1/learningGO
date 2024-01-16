package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//numeros inteiros

	fmt.Print(1, 2, 1000)
	fmt.Println("Literal inteiro é:", reflect.TypeOf(32000))
	//sem sinal(só positivos) ..... uint8 unint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal.... int 8 int 16 int32 int64
	i1 := math.MaxInt64

	fmt.Println("O valor maximo do int é", i1)
	fmt.Println("o tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela unicode(int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))

	var s1 string = `aaaa
	aa`

	fmt.Println("O rune é", reflect.TypeOf(s1))
	fmt.Println("O tamanho", len(s1))

	var char = '3'
	fmt.Println(char)
}
