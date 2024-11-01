# Simple CPU Emulator

This is a simple CPU emulator with 12 opcodes that prints the Fibonacci sequence up to n.

### Opcodes

| Instruction | Opcode | Syntax                     | Operation                     | Description                                                                 |
|-------------|--------|----------------------------|-------------------------------|-----------------------------------------------------------------------------|
| MOVR        | 10     | MOVR reg_dst, reg_src      | regdst = regsrc               | Loads the value from regsrc into regdst.                                   |
| MOVV        | 11     | MOVV reg_dst, value        | regdst = value                | Loads the numeric value into register regdst.                              |
| ADD         | 20     | ADD reg_dst, reg_src       | regdst += regsrc              | Adds the value from regsrc to the value of regdst and store the result in reg_dst. |
| SUB         | 21     | SUB reg_dst, reg_src       | regdst -= regsrc              | Subtracts the value of regsrc from the value of regdst and store the result in reg_dst. |
| PUSH        | 30     | PUSH reg_src                | stack.push(regsrc)            | Pushes the value of reg_src on the stack.                                  |
| POP         | 31     | POP reg_dst                 | regdst = stack.pop()          | Pops the last value from stack and loads it into register reg_dst.         |
| JP          | 40     | JP addr                     | PC = addr                     | Jumps the execution to address addr. Similar to a GOTO!                   |
| JL          | 41     | JL reg_1, reg_2, addr      | if (reg1 < reg2) PC = addr   | Jump to the address addr only if the value from reg1 < reg2.              |
| CALL        | 42     | CALL addr                   | stack.push(PC+1); PC = addr  | Pushes onto the stack the address of instruction that follows CALL and then jumps to address addr. |
| RET         | 50     | RET                         | PC = stack.pop()              | Pops from the stack the last number, assumes is an address and jump to that address. |
| PRINT       | 60     | PRINT reg                   | print(reg)                    | Print on the screen the value contained in the register reg.               |
| HALT        | 255    | HALT                        | stop execution                | Stops our VM. The virtual CPU doesn't execute instructions once HALT is encountered. |

### Golang Code
```
func printFibo(n int) {
    pre1 := 0
    pre2 := 1
    curr := 1

    fmt.Println(pre2)

    for curr < n {
        next := pre1 + pre2
        fmt.Println(next)
        pre1 = pre2
        pre2 = next
        curr++
    }
}
```

### Assembly Code

```
// Loads value 10 in R0 
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
```

### Output
```
This program prints Fibonacci numbersby running a machine code program on top of a VM/Virtual CPU

Tokens: [[MOVV R0 10] [CALL 6] [HALT] [PUSH R0] [MOVV R0 0] [MOVV R1 1] [MOVV R3 1] [PRINT R1] [MOV
R R2 R0] [ADD R2 R1] [PRINT R2] [MOVR R0 R1] [MOVR R1 R2] [MOVV R2 1] [ADD R3 R2] [POP R2] [PUSH R2
] [JL R3 R2 19] [POP R0] [RET]]

Bytecode: [11 0 10 42 6 255 30 0 11 0 0 11 1 1 11 3 1 60 1 10 2 0 20 2 1 60 2 10 0 1 10 1 2 11 2 1 
20 3 2 31 2 30 2 41 3 2 19 31 0 50]

Fibonacci for 10:
1
1
2
3
5
8
13
21
34
55

Disassembler:
PC      Byte Instr      ASM
0       11 0 10         MOVV R0, 10
3       42 6            CALL 6
5       255             HALT 
6       30 0            PUSH R0
8       11 0 0          MOVV R0, 0
11      11 1 1          MOVV R1, 1
14      11 3 1          MOVV R3, 1
17      60 1            PRINT R1
19      10 2 0          MOVR R2, R0
22      20 2 1          ADD R2, R1
25      60 2            PRINT R2
27      10 0 1          MOVR R0, R1
30      10 1 2          MOVR R1, R2
33      11 2 1          MOVV R2, 1
36      20 3 2          ADD R3, R2
39      31 2            POP R2
41      30 2            PUSH R2
43      41 3 2 19       JL R3R219
47      31 0            POP R0
49      50              RET 
```
