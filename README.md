# Golang Game of Life

This is a game of life simulation based on [Conways Game Of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) developed in [golang](https://golang.org/). 

![Screencapture GIF](https://thumbs.gfycat.com/CompassionateFlusteredBeardedcollie-size_restricted.gif)

## To Install

`go get github.com/poixen/go_game_of_life`

## To Run

`go run {workspace}/github.com/poixen/go_game_of_life/gol.go`

## Additional Information

Simple program to helpo learn `golang`. Uses go routines, channels, pointers, structs, pointer-receivers, mutex locking and wait groups.

Average time to perform a step is ~`4ms`.

## Known Issues

The number of `Alive` Cells are more than what are visible. This is most likely due to the game exceeding the limits of the visible board. This can also be the reason why cells can become alive from the edge of the board with no surrounding activity.