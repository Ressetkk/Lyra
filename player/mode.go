package player

import (
	"github.com/Ressetkk/lyra/score"
	"strconv"
)

type ModeFunc func(cmd string) func(out chan score.Note, in *score.Note)

var modes = map[rune]ModeFunc{
	'd': DelayMode,
}

// TODO asynchronous mode handling in key event
func DelayMode(cmd string) func(out chan score.Note, in *score.Note) {
	_, err := strconv.Atoi(cmd)
	if err != nil {
		panic(err)
	}
	return func(out chan score.Note, in *score.Note) {
		for _, n := range in.Notes {
			out <- score.Note{Notes: []int{n}}
		}
	}
}
