package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var instructions = map[string]int{
	"MOVR":  MOVR,
	"MOVV":  MOVV,
	"ADD":   ADD,
	"SUB":   SUB,
	"PUSH":  PUSH,
	"POP":   POP,
	"JP":    JP,
	"JL":    JL,
	"CALL":  CALL,
	"RET":   RET,
	"PRINT": PRINT,
	"HALT":  HALT,
}

var registerNames = map[string]int{
	"R0": R0,
	"R1": R1,
	"R2": R2,
	"R3": R3,
}

func Assemble(code string) []int {
	tokens := getTokens(code)
	fmt.Printf("tokens: %+v\n", tokens)
	bytes := getBytecode(tokens)
	fmt.Printf("bytes: %+v\n", bytes)
	return bytes
}

func getTokens(code string) [][]string {
	re := regexp.MustCompile(`\r?\n`)
	arLines := re.Split(code, -1)
	var tokens [][]string

	for _, line := range arLines {
		txt := strings.TrimSpace(line)
		if txt == "" || strings.HasPrefix(txt, "//") {
			continue
		}
		re := regexp.MustCompile(`[\s,]+`)
		spTxt := re.Split(txt, -1)

		tokens = append(tokens, spTxt)
	}

	return tokens
}

func getBytecode(tokens [][]string) []int {
	var bytes []int

	for _, line := range tokens {
		for i, tkn := range line {
			token := strings.ToUpper(strings.TrimSpace(tkn))
			if i == 0 {
				bytes = append(bytes, instructions[token])
			} else {
				if strings.HasPrefix(token, "R") {
					bytes = append(bytes, registerNames[token])
					continue
				}
				atoi, err := strconv.Atoi(token)
				if err != nil {
					panic(err)
				}
				bytes = append(bytes, atoi)
			}
		}
	}

	return bytes
}
