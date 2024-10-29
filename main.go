package main

import (
	"fmt"
)

func main() {
	fmt.Printf("This program prints Fibonacci numbers")
	fmt.Printf("by running a machine code program on top of a VM/Virtual CPU\n")
	code := `// Loads value 10 in R0 
// and calls Fibonacci routine

MOVV R0, 10
CALL 6
HALT

// This is the Fibonacci routing
// Expects number of Fibonacci 
// numbers in register R0

PUSH R0
MOVV R0, 0
MOVV R1, 1
MOVV R3, 1
PRINT R1
MOVR R2, R0
ADD R2, R1
PRINT R2
MOVR R0, R1
MOVR R1, R2
MOVV R2, 1
ADD R3, R2
POP R2
PUSH R2
JL R3, R2, 19
POP R0
RET
`

	bytecode := Assemble(code)
	RunCPU(bytecode)
	Disassemble(bytecode)
}
