// https://www.urionlinejudge.com.br/judge/pt/problems/view/1041

package main

import (
	"fmt"
)

func main() {
	var x, y float64
	fmt.Scanf("%f", &x)
	fmt.Scanf("%f", &y)

	switch {
	case y == 0.0 && x == 0.0:
		fmt.Println("Origem")
	case x == 0.0:
		fmt.Println("Eixo Y")
	case y == 0.0:
		fmt.Println("Eixo X")
	case y > 0.0 && x > 0.0:
		fmt.Println("Q1")
	case y > 0.0 && x < 0.0:
		fmt.Println("Q2")
	case y < 0.0 && x < 0.0:
		fmt.Println("Q3")
	case y < 0.0 && x > 0.0:
		fmt.Println("Q4")
	}
}
