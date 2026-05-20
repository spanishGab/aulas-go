package basics

import (
	"fmt"
	"math"
	"strings"
)

func Integers() {
	println("INTEGERS:")
	// O valor após o nome do tipo indica quantos bits são utilizados para armazenar o valor.
	// Quanto mais bits, maior o intervalo de valores que podem ser representados.
	// int8, int16, int32, int64
	// uint8, uint16, uint32, uint64

	// int // permite que o compilador escolha quantos bits serão utilizados
	// uint // permite que o compilador escolha quantos bits serão utilizados
	// byte // alias para uint8. Utilizado para enfatizar que o valor é um dado bruto, não um número
	// rune // alias para int32. Representa um ponto de código Unicode

	var a int = 10
	var b uint = 20
	var c byte = 10
	var d rune = 'A'

	fmt.Printf("\ta=%d; b=%d; (c=%08[3]b OR c=%[3]d); (d=%[4]d OR d=%[4]c OR d=%[4]q)\n", a, b, c, d)

	var smallInt int8 = math.MaxInt8
	var smallUint uint8 = math.MaxUint8
	fmt.Printf("\tCUIDADO - integer overflow: %d\n", smallInt+1)
	fmt.Printf("\tCUIDADO - unsigned integer overflow: %d\n", smallUint+1)
	println()
}

func Floats() {
	println("FLOATS:")
	// float32, float64 (prefira float64, pois permite menos erros)
	var a float32 = math.Phi
	var b float64 = math.Pi

	fmt.Printf("\t(a=%[1]f OR a=%[1]g); (b=%[2]f OR b=%[2]g)\n", a, b)
	println()
}

func Booleans() {
	println("BOOLEANS:")
	// bool
	var a bool = true
	var b bool = false

	// var c int = 1
	// if c { // valor deve ser um booleano
	// }
	// for a { // valor deve ser um booleano
	// }

	fmt.Printf("\ta=%t; b=%t\n", a, b)
	println()
}

func Strings() {
	println("STRINGS:")

	// string é uma sequência de runes (unicode codificado em UTF-8), ou seja, cada caractere ocupa 32 bits.
	var a string = "Hello, World!"

	// strings são imutáveis
	b := a                    // não é uma cópia do valor de a, é apenas uma nova referência para a mesma string
	a += " I am learning Go." // porém uma VARIÁVEL do tipo string é mutável (já não referencia a mesma string que b)
	fmt.Printf("\ta=%s; b=%s\n", a, b)

	// a[0] = '.' // erro: strings são imutáveis

	sliced := a[13:19]  // apenas referencia a string original, sem criar uma nova string
	welcome := a[:13]   // apenas referencia a string original, sem criar uma nova string
	education := a[13:] // apenas referencia a string original, sem criar uma nova string

	fmt.Printf("\tsliced=%q; welcome=%q; education=%q\n", strings.Trim(sliced, " "), welcome, education)

	rawLiteral := `{
	  "id": 1,
	  "name": "basic types",
	  "description": "basic types in Go"
	}`
	fmt.Printf("\trawLiteral=%s\n", rawLiteral)
	println()
}
