package main

import (
	"fmt"
	"math"
)

func main() {

	var objetivo int64

	epsilon := 0.001
	bajo := 0.0

	fmt.Println("Escoga un número:")
	fmt.Scanf("%d", &objetivo)

	alto := math.Max(1.0, float64(objetivo))

	respuesta := (alto + bajo) / 2

	for math.Abs(math.Pow(respuesta, 2) - float64(objetivo)) >= epsilon {
		fmt.Printf("bajo es %v, alto es %v y la respuesta es %v\n", bajo, alto, respuesta)

		if math.Pow(respuesta, 2) < float64(objetivo) {
			bajo = respuesta
		} else {
			alto = respuesta
		}

		respuesta = (alto + bajo) / 2
	}

	fmt.Printf("La raiz cuadrada de %v es %v\n", objetivo, respuesta)

}
