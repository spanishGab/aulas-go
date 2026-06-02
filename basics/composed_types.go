package basics

import (
	"fmt"
	"strings"
	"time"
)

func Arrays() {
	fmt.Println("ARRAYS:")

	// Arrays são imutáveis e possuem tamanho fixo.
	// [size]type
	// O tamanho do array é parte do seu tipo. Arrays de tamanhos diferentes são tipos diferentes.
	var arr1 [3]int = [3]int{1, 2, 3}
	arr2 := arr1 // cria uma cópia de arr1 e atribui a arr2
	arr1[0] = -1

	// Em NodeJS:
	// > let arr1 = [1,2,3]
	// > arr1
	// > [ 1, 2, 3 ]
	// > let arr2 = arr1
	// > arr2
	// > [ 1, 2, 3 ]
	// > arr1[0] = -1
	// > -1
	// > arr1
	// > [ -1, 2, 3 ]
	// > arr2
	// > [ -1, 2, 3 ]

	var equals bool = arr1 == arr2 // comparação de ARRAYS é possível utilizando diretamente os operadores == ou !=

	bigArray := [...]int{0: 7, 99: 42}

	fmt.Printf("\tarr1=%v; arr2=%v, equals=%t\n", arr1, arr2, equals)
	fmt.Printf("\tbigArray=%v\n", bigArray)
	println()
}

func Slices() {
	println("SLICES:")
	// um slice é uma referência a um array subjacente. Ele é mutável e tem um tamanho dinâmico.
	// []type

	arr := [5]string{"go", "is", "funny", "and", "powerful"}
	slice1 := arr[0:3] // cria um slice que referencia os elementos do índice 0 ao 2 do array
	slice2 := slice1   // slice2 referencia o mesmo array subjacente que slice1

	arr[0] = "golang"
	fmt.Printf("\tQuantidade de itens em slice1: %d; Capacidade de slice1: %d\n", len(slice1), cap(slice1))
	fmt.Printf("\tslice1=%v; slice2=%v\n", slice1, slice2)

	fmt.Printf("\tJoinned slice2: %s\n", strings.Join(slice2, ","))

	for i, word := range slice1 {
		fmt.Printf("\tindex=%d, word=%s\n", i, word)
	}

	var slice3 = make([]int, 3) // cria um slice de strings com comprimento 3 e capacidade 3
	fmt.Printf("\tslice3=%v", slice3)
	println()
}

func Maps() {
	println("MAPS:")
	// um map é uma referência a uma tabela hash. Ele é mutável e tem um tamanho dinâmico.
	// map[keyType]valueType

	ages := map[string]int{
		"Luna": 30,
		"Ana":  25,
		"João": 20,
	}
	fmt.Printf("\tQuantidade de chaves em ages: %d\n", len(ages))
	fmt.Printf("\tages=%v\n", ages)

	println("\tAdicionando chave 'Maria'")
	ages["Maria"] = 32
	fmt.Printf("\tages=%v\n", ages)

	println("\tRemovendo chave 'Maria'")
	delete(ages, "Maria")
	fmt.Printf("\tages=%v\n", ages)

	println("\tIterando sobre ages")
	for name, age := range ages {
		fmt.Printf("\tname=%s, age=%d\n", name, age)
	}

	println("\tVerificando se chave 'Luna' existe")
	if lunaAge, ok := ages["Luna"]; ok {
		fmt.Printf("\tLuna tem %d anos\n", lunaAge)
	} else {
		fmt.Println("\tChave 'Luna' não existe")
	}
	println()
}

func Structs() {
	// Uma struct é um tipo de dado composto que agrega multiplos valores de tipos diferentes
	type Animal struct {
		Breed       string
		Name        string
		Age         int
		IsCastrated bool
		Vaccines    map[string]time.Time
	}

	animal := Animal{
		Breed:       "Shihtzu",
		Name:        "Thor",
		Age:         14,
		IsCastrated: true,
		Vaccines: map[string]time.Time{
			"rage": time.Now(),
		},
	}
	fmt.Printf("animal=%#v", animal)

	thor := &animal

	animal.Breed = "No Breed"
	animal.Name = "Safira"
	fmt.Printf("thor.Name=%s, animal.Name=%s", thor.Name, animal.Name)

}
