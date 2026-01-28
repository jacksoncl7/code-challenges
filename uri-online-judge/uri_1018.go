// https://www.urionlinejudge.com.br/judge/pt/problems/view/1018
package main

import "fmt"

func print_notes(count int, str_notes string) {
	fmt.Printf("%v nota(s) de R$ %v\n", count, str_notes)
}

func counter(note int, value int) int {
	count := 0
	for note <= value {
		value -= note
		count++
	}
	return count
}

func main() {
	value := 0
	fmt.Scan(&value)
	fmt.Println(value)
	hundred := counter(100, value)
	partial_value := value - 100 * hundred
	fifty := counter(50, partial_value)
	partial_value =	partial_value - 50 * fifty
	twenty := counter(20, partial_value)
	partial_value =	partial_value - 20 * twenty
	ten := counter(10, partial_value)
	partial_value =	partial_value - 10 * ten
	five := counter(5, partial_value)
	partial_value =	partial_value - 5 * five
	two := counter(2, partial_value)
	partial_value =	partial_value - 2 * two
	one := counter(1, partial_value)

	print_notes(hundred, "100,00")
	print_notes(fifty, "50,00")
	print_notes(twenty, "20,00")
	print_notes(ten, "10,00")
	print_notes(five, "5,00")
	print_notes(two, "2,00")
	print_notes(one, "1,00")
}
