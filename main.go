package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func loadMaze(filename string) ([]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	data := make([]string, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data, nil
}

func globalizeLevels(levels []string) error {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return errors.New("Unable to retrive path to the current source file")
	}
	dir, err := filepath.Abs(filepath.Dir(filename))
	if err != nil {
		return err
	}

	for i := range levels {
		levels[i] = filepath.Join(dir, levels[i])
	}

	return nil
}

func printMaze(maze []string) {
	clearScreen()
	for _, row := range maze {
		fmt.Println(row)
	}
}

func init() {
	cbTerm := exec.Command("/bin/stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()

	if err != nil {
		log.Fatalln("Unable to activate cbreak mode terminal: %v\n", err)
	}
}

func cleanup() {
	cookedTerm := exec.Command("/bin/stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatalln("Unable to enable cooked mode: %v", err)
	}
}

func readInput() (string, error) {
	buffer := make([]byte, 100)

	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	}
	return "", nil
}

func clearScreen() {
	fmt.Print("\x1b[2J")
	moveCursor(0, 0)
}

func moveCursor(row, col int) {
	fmt.Printf("\x1b[%d;%df", row+1, col+1)
}

func main() {
	defer cleanup()
	// initialize game
	levels := []string{filepath.Join("step01", "maze01.txt")}
	err := globalizeLevels(levels)
	if err != nil {
		log.Fatalln(err)
	}
	maze, err := loadMaze(levels[0])
	if err != nil {
		log.Fatalln("Unable to find the level file: \n", err)
	}
	// load resources

	// game loop
	for {
		// update screen
		printMaze(maze)
		// process input
		input, err := readInput()
		if err != nil {
			log.Fatalln("Unable to process keyboard input:\n%v", err)
		}
		if input == "ESC" {
			break
		}
		// process movement

		// process collisions

		// check game over

		// Temp: break infinite loop

		// repeat
	}
}
