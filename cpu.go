package main

import "fmt"

var (
	registers = []int{0, 0, 0, 0}
	pc        = 0
	halted    = false
	program   = []int{}
	stack     = []int{}
)

var (
	MOVR  = 10
	MOVV  = 11
	ADD   = 20
	SUB   = 21
	PUSH  = 30
	POP   = 31
	JP    = 40
	JL    = 41
	CALL  = 42
	RET   = 50
	PRINT = 60
	HALT  = 255
)

var (
	R0 = 0
	R1 = 1
	R2 = 2
	R3 = 3
)

func runInstr() {
	if halted {
		return
	}
	instr := program[pc]

	switch instr {
	case MOVR:
		pc++
		rdst := program[pc]
		pc++
		rsrc := program[pc]
		pc++
		registers[rdst] = registers[rsrc]
	case MOVV:
		pc++
		rdst := program[pc]
		pc++
		val := program[pc]
		pc++
		registers[rdst] = val
	case ADD:
		pc++
		rdst := program[pc]
		pc++
		rsrc := program[pc]
		pc++
		registers[rdst] += registers[rsrc]
	case SUB:
		pc++
		rdst := program[pc]
		pc++
		rsrc := program[pc]
		pc++
		registers[rdst] -= registers[rsrc]
	case PUSH:
		pc++
		rsrc := program[pc]
		pc++
		stack = append(stack, registers[rsrc])
	case POP:
		pc++
		rdst := program[pc]
		pc++
		registers[rdst], stack = stack[len(stack)-1], stack[:len(stack)-1]
	case JP:
		pc++
		addr := program[pc]
		pc++
		pc = addr
	case JL:
		pc++
		reg1 := program[pc]
		pc++
		reg2 := program[pc]
		pc++
		addr := program[pc]
		pc++
		if registers[reg1] < registers[reg2] {
			pc = addr
		}
	case CALL:
		pc++
		addr := program[pc]
		pc++ // maybe
		stack = append(stack, pc)
		pc = addr
	case RET:
		pc++
		pc, stack = stack[len(stack)-1], stack[:len(stack)-1]
	case PRINT:
		pc++
		reg := program[pc]
		pc++
		fmt.Printf("%d\n", registers[reg])
	case HALT:
		pc++
		halted = true
	}
}

func RunCPU(bytecode []int) {
	program = bytecode
	for !halted {
		runInstr()
	}
}
