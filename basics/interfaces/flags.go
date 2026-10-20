package interfaces

import (
	"flag"
	"fmt"
	"strings"
)

type Kilograms float64

func (k Kilograms) String() string {
	return fmt.Sprintf("%.2gKg", k)
}

type kilogramsFlag struct {
	Kilograms
}

func (k *kilogramsFlag) String() string {
	return k.Kilograms.String()
}

func (k *kilogramsFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "Kg":
		k.Kilograms = Kilograms(value)
		return nil
	case "g":
		k.Kilograms = Kilograms(value / 1000)
		return nil
	default:
		return fmt.Errorf("invalid weight unit %q", s)
	}
}

func KilogramsFlag(name string, value Kilograms, usage string) *Kilograms {
	k := kilogramsFlag{value}
	flag.CommandLine.Var(&k, name, usage)
	return &k.Kilograms
}

var kilos = KilogramsFlag("kilos", 5.0, "the weight in Kilograms")
var gramas = flag.Int("gramas", 5, "Peso em gramas")

func ConvertToKilograms() {
	flag.Parse()
	fmt.Printf("%s has %s", strings.Join(flag.Args(), ""), *kilos)
}
