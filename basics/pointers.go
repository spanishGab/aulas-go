package basics

import "fmt"

func Pointers() {
	println("POINTERS:")
	// uma VARIÁVEL é um item de armazenagem que contém um valor.
	// um PONTEIRO é o endereço de uma variável

	a := 10
	var p *int = &a

	fmt.Printf("\ta=%d; &a=%p == p=%p\n", a, &a, p)

	increment(&a)
	fmt.Printf("\tDepois de increment(): a=%d; &a=%p == p=%p\n", a, &a, p)

	// o elemento de um array é essencialemente como uma variável, portanto podemos acessar seu endereço
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Printf("\tendereços dos elementos de arr: &arr[0]=%p; &arr[1]=%p; &arr[2]=%p\n", &arr[0], &arr[1], &arr[2])
	println()
}

func increment(x *int) {
	*x++
}
