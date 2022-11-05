//O objetivo deste programa é converter o valor do ponto de ebulição da água de Kelvin para Celsius//

package main

import "fmt"

const ebulicaoK = 373.0 //Temperatura de ebulição da água em Kelvin//

func main() {

	K := ebulicaoK //K representa Kelvin//

	C := (K - 273) //Utilizei a fórmula que foi passada para encontrar o valor da temperatura de ebulição em Celsius. C representa Celsius//

	fmt.Printf("Ponto de ebulição da água em °K é: %g e o ponto de ebulição da água em °C é: %g ", K, C)
}
