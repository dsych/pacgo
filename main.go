package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
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
	for _, row := range maze {
		fmt.Println(row)
	}
}

func main() {
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

		// process movement

		// process collisions

		// check game over

		// Temp: break infinite loop
		fmt.Println("Hello, Pac Go!")
		break

		// repeat
	}
}
