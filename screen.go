package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func SetTerm(isCooked bool) {
	var term *exec.Cmd
	if isCooked {
		term = exec.Command("stty", "-cbreak", "echo")
	} else {
		term = exec.Command("stty", "cbreak", "-echo")
	}
	term.Stdin = os.Stdin

	err := term.Run()
	if err != nil {
		log.Fatalln("Unable to enable cooked mode: %v", err)
	}
}

func ReadInput() (bool, []int, error) {
	buffer := make([]byte, 100)

	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return false, nil, err
	}
	// 0 - y
	// 1 - x
	coords := []int{0, 0}

	if cnt == 1 && buffer[0] == 0x1b {
		return true, nil, nil
	} else if cnt >= 3 && buffer[0] == 0x1b && buffer[1] == '[' {
		switch buffer[2] {
		case 'A':
			coords[0] = -1
		case 'B':
			coords[0] = 1
		case 'C':
			coords[1] = 1
		case 'D':
			coords[1] = -1
		}
	}
	return false, coords, nil
}

func ClearScreen() {
	fmt.Print("\x1b[2J")
	moveCursor(0, 0)
}

func moveCursor(row, col int) {
	fmt.Printf("\x1b[%d;%df", row+1, col+1)
}
