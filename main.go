package main

import "aulas/basics"

func main() {
	commaSep := &basics.CommaSeparator{}
	dotSep := &basics.DotSeparator{}
	// pipeSep := &basics.PipeSeparator{}

	basics.HandleString("aprender golang é bem legal com vírgulas", commaSep)
	basics.HandleString("aprender golang é bem legal com pontos", dotSep)
	// basics.HandleString("aprender golang é bem legal com pontos", pipeSep)
}
