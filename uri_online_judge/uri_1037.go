// https://www.urionlinejudge.com.br/judge/pt/problems/view/1037

package main

import "fmt"

func main() {
	var indata float64
	var outdata string
	fmt.Scanf("%f", &indata)

	switch {
	case indata >= 0 && indata <= 25:
		outdata = "Intervalo [0,25]"
	case indata > 25 && indata <= 50:
		outdata = "Intervalo (25,50]"
	case indata > 50 && indata <= 75:
		outdata = "Intervalo (50,75]"
	case indata > 75 && indata <= 100:
		outdata = "Intervalo (75,100]"
	default:
		outdata = "Fora de intervalo"
	}
	fmt.Printf("%s\n", outdata)
}
