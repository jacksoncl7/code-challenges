// https://www.urionlinejudge.com.br/judge/pt/problems/view/1002

package main

import ( "fmt" )

func main() {
	PI := 3.14159
	var area float64
	var raio float64

	fmt.Scanf("%f", &raio)
	area = raio*raio*PI
	fmt.Printf("A=%.4f\n", area)
}
