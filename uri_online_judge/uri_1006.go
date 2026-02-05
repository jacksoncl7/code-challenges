// https://www.urionlinejudge.com.br/judge/pt/problems/view/1006

package main

import "fmt"

func main() {
	var first, second, third float64
	fmt.Scanf("%f", &first)
	fmt.Scanf("%f", &second)
	fmt.Scanf("%f", &third)

	media := (2.0*first+3.0*second+5.0*third)/10.0
	fmt.Printf("MEDIA = %.1f\n", media)
}
