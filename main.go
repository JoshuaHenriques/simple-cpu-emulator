package main

var (
	registers = []int{0, 0, 0, 0}
	pc        = 0
	halted    = false
	program   = []int{
		11, 0, 10, 42, 6, 255, 30, 0, 11, 0, 0, 11, 1, 1, 11, 3, 1, 60, 1, 10, 2, 0, 20,
		2, 1, 60, 2, 10, 0, 1, 10, 1, 2, 11, 2, 1, 20, 3, 2, 31, 2, 30, 2, 41, 3, 2, 19, 31, 0, 50,
	}
)

func runInstr() {
	if halted {
		return
	}
	instr := program[pc]

	switch instr {
	//
	}
}

func main() {
	for !halted {
		runInstr()
	}
}