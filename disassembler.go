package main

import (
	"fmt"
	"strconv"
	"strings"
)

var regs = []string{"R0", "R1", "R2", "R3"}

func getAssembly(pc int, bytes []int, txt string) string {
	src := strings.Builder{}
	src.WriteString(strconv.Itoa(pc))
	src.WriteString("\t")

	for _, byte := range bytes {
		src.WriteString(strconv.Itoa(byte))
		src.WriteString(" ")
	}

	currLen := src.Len()
	if currLen < 15 {
		padding := strings.Repeat(" ", 15-src.Len())
		src.WriteString(padding)
	}
	src.WriteString("\t")
	src.WriteString(txt)
	src.WriteString("\n")

	return src.String()
}

func Disassemble(program []int) string {
	pc := 0
	src := &strings.Builder{}

	for pc < len(program) {
		currPC := pc
		instr := program[pc]
		instruction := &strings.Builder{}

		switch instr {
		case MOVR:
			pc++
			rdst := program[pc]
			pc++
			rsrc := program[pc]
			pc++
			instruction.Reset()
			instruction.WriteString("MOVR ")
			instruction.WriteString(regs[rdst])
			instruction.WriteString(", ")
			instruction.WriteString(regs[rsrc])
			src.WriteString(getAssembly(currPC, []int{instr, rdst, rsrc}, instruction.String()))
		case MOVV:
			pc++
			rdst := program[pc]
			pc++
			val := program[pc]
			pc++
			instruction.Reset()
			instruction.WriteString("MOVV ")
			instruction.WriteString(regs[rdst])
			instruction.WriteString(", ")
			instruction.WriteString(strconv.Itoa(val))
			src.WriteString(getAssembly(currPC, []int{instr, rdst, val}, instruction.String()))
		case ADD:
			pc++
			rdst := program[pc]
			pc++
			rsrc := program[pc]
			pc++
			instruction.Reset()
			instruction.WriteString("ADD ")
			instruction.WriteString(regs[rdst])
			instruction.WriteString(", ")
			instruction.WriteString(regs[rsrc])
			src.WriteString(getAssembly(currPC, []int{instr, rdst, rsrc}, instruction.String()))
		case SUB:
			pc++
			rdst := program[pc]
			pc++
			rsrc := program[pc]
			pc++
			instruction.Reset()
			instruction.WriteString("SUB ")
			instruction.WriteString(regs[rdst])
			instruction.WriteString(", ")
			instruction.WriteString(regs[rsrc])
			src.WriteString(getAssembly(currPC, []int{instr, rdst, rsrc}, instruction.String()))
		case PUSH:
			pc++
			rsrc := program[pc]
			pc++
			instruction.Reset()
			instruction.WriteString("PUSH ")
			instruction.WriteString(regs[rsrc])
			src.WriteString(getAssembly(currPC, []int{instr, rsrc}, instruction.String()))
		case POP:
			pc++
			rdst := program[pc]
			pc++
			instruction.Reset()
			instruction.WriteString("POP ")
			instruction.WriteString(regs[rdst])
			src.WriteString(getAssembly(currPC, []int{instr, rdst}, instruction.String()))
		case JP:
			pc++
			addr := program[pc]
			pc++
			instruction.Reset()
			instruction.WriteString("JP ")
			instruction.WriteString(strconv.Itoa(addr))
			src.WriteString(getAssembly(currPC, []int{instr, addr}, instruction.String()))
		case JL:
			pc++
			reg1 := program[pc]
			pc++
			reg2 := program[pc]
			pc++
			addr := program[pc]
			pc++
			instruction.Reset()
			instruction.WriteString("JL ")
			instruction.WriteString(regs[reg1])
			instruction.WriteString(regs[reg2])
			instruction.WriteString(strconv.Itoa(addr))
			src.WriteString(getAssembly(currPC, []int{instr, reg1, reg2, addr}, instruction.String()))
		case CALL:
			pc++
			addr := program[pc]
			pc++
			instruction.Reset()
			instruction.WriteString("CALL ")
			instruction.WriteString(strconv.Itoa(addr))
			src.WriteString(getAssembly(currPC, []int{instr, addr}, instruction.String()))
		case RET:
			pc++
			instruction.Reset()
			instruction.WriteString("RET ")
			src.WriteString(getAssembly(currPC, []int{instr}, instruction.String()))
		case PRINT:
			pc++
			reg := program[pc]
			pc++
			instruction.Reset()
			instruction.WriteString("PRINT ")
			instruction.WriteString(regs[reg])
			src.WriteString(getAssembly(currPC, []int{instr, reg}, instruction.String()))
		case HALT:
			pc++
			instruction.Reset()
			instruction.WriteString("HALT ")
			src.WriteString(getAssembly(currPC, []int{instr}, instruction.String()))
		default:
			pc = len(program)
			instruction.Reset()
			instruction.WriteString("UNK - disassemble stopped")
			src.WriteString(getAssembly(currPC, []int{instr}, instruction.String()))
		}
	}

	fmt.Printf("\nDisassembler:")
	fmt.Printf("\n%-8s%-16s%-s\n", "PC", "Byte Instr", "ASM")
	fmt.Printf("%s", src.String())
	return src.String()
}
