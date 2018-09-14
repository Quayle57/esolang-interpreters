package main

import (
	"fmt"
)

func moveRight(code, tape *string, cpos, tpos *int) {
	*tpos++
}

func moveLeft(code, tape *string, cpos, tpos *int) {
	*tpos--
}

func flip(code, tape *string, cpos, tpos *int) {
	if byte((*tape)[*tpos]) == byte('0') {
		*tape = (*tape)[:*tpos] + string(byte('1')) + (*tape)[*tpos+1:]
	} else {
		*tape = (*tape)[:*tpos] + string(byte('0')) + (*tape)[*tpos+1:]
	}
}

func jumpRight(code, tape *string, cpos, tpos *int) {
	toSkip := 0
	vMax := len(*code)

	if (*tape)[*tpos] == byte('0') {
		for *cpos < vMax {
			*cpos++
			if (*code)[*cpos] == '[' {
				toSkip++
			} else if (*code)[*cpos] == ']' && toSkip > 0 {
				toSkip--
			} else if (*code)[*cpos] == ']' && toSkip == 0 {
				break
			}
		}
	}
}

func jumpLeft(code, tape *string, cpos, tpos *int) {
	toSkip := 0

	if (*tape)[*tpos] == byte('1') {
		for *cpos > 0 {
			*cpos--
			if (*code)[*cpos] == ']' {
				toSkip++
			} else if (*code)[*cpos] == '[' && toSkip > 0 {
				toSkip--
			} else if (*code)[*cpos] == '[' && toSkip == 0 {
				break
			}
		}
	}
}

type handle struct {
	c uint8
	f func(code, tape *string, cpos, tpos *int)
}

func interpreter(code, tape string) string {
	tpos, cpos := 0, 0
	max := len(code)
	// https://esolangs.org/wiki/Smallfuck
	tab := []handle{
		handle{'>', moveRight},
		handle{'<', moveLeft},
		handle{'*', flip},
		handle{'[', jumpRight},
		handle{']', jumpLeft},
	}

	for cpos < max {
		c := code[cpos]
		// check if the command is known and trigger it if it is
		for _, member := range tab {
			if member.c == uint8(c) {
				member.f(&code, &tape, &cpos, &tpos)
			}
		}
		cpos++
		// check if out of bound
		if tpos >= len(tape) || tpos < 0 {
			return tape
		}
	}
	return tape
}

func main() {
	code := "*[>*]"
	tape := "000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	fmt.Println("Program", code, "\nTape before", tape, "\nThe program should flip every bit from the tape to 1 :\nTape after", interpreter(code, tape))
}
