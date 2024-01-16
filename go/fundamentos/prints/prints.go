package main

import "fmt"

func main() {

	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha")

	x := 3.141516

	fmt.Println("O valor de x é", x) // permite concatenar com variaveis

	xs := fmt.Sprint(x) // converte para string uma variavel

	fmt.Println("O valor de xs é " + xs) // mão mostra espaço no print da variavel

	fmt.Printf("O valor de x é %f\n", x)  // imprimir sem formatação
	fmt.Printf("O valor de x é %.2f.", x) //imprimir com formatação

	a, b, c, d := 1, 1.9999, false, "Opa"

	fmt.Printf("\n %d \n %f %.1f \n %t \n %s \n", a, b, b, c, d)
	fmt.Printf("\n %v \n %v %.v \n %v \n %v \n", a, b, b, c, d)
}
