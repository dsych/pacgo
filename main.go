package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
)

func loadMaze(filename string) ([]string, Player, []*Player, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, Player{}, nil, err
	}

	defer file.Close()

	data := make([]string, 0)

	scanner := bufio.NewScanner(file)

	var player Player
	var ghosts []*Player

	found := false

	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	for row, line := range data {
		for col, chr := range line {
			if chr == 'P' {
				player = Player{x: col, y: row}
				found = true
			} else if chr == 'G' {
				ghosts = append(ghosts, &Player{x: col, y: row})
			}
		}
	}
	if !found {
		return nil, Player{}, nil, errors.New("Unable to locate player on the board")
	}

	return data, player, ghosts, nil
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

func printMaze(maze []string, p *Player, ghosts []*Player) {
	ClearScreen()
	for _, row := range maze {
		for _, chr := range row {
			switch chr {
			case '#':
				fmt.Printf("%c", chr)
			default:
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
	moveCursor(p.y, p.x)
	fmt.Printf("P")

	for _, g := range ghosts {
		moveCursor(g.y, g.x)
		fmt.Printf("G")
	}
}

func movePlayer(maze []string, newCoords []int, p *Player) {
	y, x := p.getCoords()
	y = (y + newCoords[0]) % len(maze)
	x = (x + newCoords[1]) % len(maze[y])
	if maze[y][x] != '#' {
		p.y = y
		p.x = x
	}
}

func moveGhosts(maze []string, ghosts []*Player) {

	for _, g := range ghosts {
		newCoords := []int{0, 0}
		switch dir := rand.Intn(4); dir {
		case 0:
			newCoords[0] = -1
		case 1:
			newCoords[0] = 1
		case 2:
			newCoords[1] = 1
		case 3:
			newCoords[1] = -1
		}
		movePlayer(maze, newCoords, g)
	}
}

func init() {
	SetTerm(false)
}

func main() {
	defer SetTerm(true)
	// initialize game
	levels := []string{filepath.Join("step01", "maze01.txt")}
	err := globalizeLevels(levels)
	if err != nil {
		log.Fatalln(err)
	}
	maze, player, ghosts, err := loadMaze(levels[0])
	if err != nil {
		log.Fatalln("Unable to find the level file: \n", err)
	}
	// load resources

	// game loop
	for {
		// update screen
		printMaze(maze, &player, ghosts)
		// process input
		isEscape, adjust, err := ReadInput()
		if err != nil {
			log.Fatalln("Unable to process keyboard input:\n%v", err)
		}
		if isEscape {
			break
		}
		// process movement
		movePlayer(maze, adjust, &player)
		moveGhosts(maze, ghosts)
		// process collisions

		// check game over

		// Temp: break infinite loop

		// repeat
	}
}
