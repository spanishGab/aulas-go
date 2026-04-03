package main

type Entries struct {
	Operation string
	Operand1  int
	Operand2  int
}

var calculatorEntries []Entries = []Entries{
	{Operation: "sum", Operand1: 2, Operand2: 3},
	{Operation: "subtract", Operand1: 5, Operand2: 2},
	{Operation: "multiply", Operand1: 3, Operand2: 4},
	{Operation: "divide", Operand1: 10, Operand2: 2},
	{Operation: "sum", Operand1: -10, Operand2: 2},
	// {Operation: "divide", Operand1: 10, Operand2: 0}, // descomente caso queira se aventurar tratando o erro
}

func main() {
	total := 0
	operationResult := 0
	for _, entry := range calculatorEntries {
		Calculate(entry.Operation, entry.Operand1, entry.Operand2, &operationResult)
		total += operationResult
	}
	println("Total:", total)
}

func Calculate(operation string, operand1 int, operand2 int, result *int) {
	switch operation {
	case "sum":
		*result = operand1 + operand2
	case "subtract":
		*result = operand1 - operand2
	case "multiply":
		*result = operand1 * operand2
	case "divide":
		if operand2 == 0 {
			println("Error: division by zero")
			return
		}
		*result = operand1 / operand2
	default:
		println("Error: unknown operation")
		return
	}
}
