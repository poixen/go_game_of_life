package main

import (
	"github.com/poixen/gol/components"
	"math/rand"
	"time"
	"fmt"
	"strconv"
	"os"
)

var BOARD_WIDTH int = 60
var BOARD_HEIGHT int = 20

func main() {
	setBoardDimensions()
	var GAME_LENGTH int = 500
	var GAME_LENGTH_STR = strconv.Itoa(GAME_LENGTH)
	var START_ALIVE int = (BOARD_WIDTH * BOARD_HEIGHT) / 2
	var BOARD_TOTAL_STR = strconv.Itoa(BOARD_WIDTH * BOARD_HEIGHT)

	rand.Seed(time.Now().UTC().UnixNano())
	for {
		board := components.NewBoard(BOARD_WIDTH, BOARD_HEIGHT)

		// set {START_ALIVE} random Cells to be alive
		for i := 0; i < START_ALIVE; i++ {
			board.SetAlive(components.NewPoint(rand.Intn(BOARD_WIDTH),rand.Intn(BOARD_HEIGHT)))
		}

		// play game for {GAME_LENGTH} frames
		for i := 0; i < GAME_LENGTH; i++ {
			now := time.Now()
			board = board.Next()
			board.Print(BOARD_WIDTH, BOARD_HEIGHT)
			fmt.Println(strconv.Itoa(i) + "/" + GAME_LENGTH_STR +
				"\tTime/Step: " + time.Since(now).String() + " \tAlive: " +
				strconv.Itoa(board.TotalAlive()) + "/" + BOARD_TOTAL_STR)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func setBoardDimensions() {
	if len(os.Args) != 3 {
		return
	}

	if w, err := strconv.Atoi(os.Args[1]); err == nil {
		BOARD_WIDTH = w
	}
	if h, err := strconv.Atoi(os.Args[2]); err == nil {
		BOARD_HEIGHT = h
	}
}
