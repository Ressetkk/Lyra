package main

import (
	"github.com/Ressetkk/windblume-lyre-player/score"
	"github.com/micmonay/keybd_event"
	"runtime"
	"time"
)

func main() {
	time.Sleep(5 *time.Second)
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}
	s := score.NewScore("a4/8 c5/8 f3a4d5/6 e5/16 d5/4 h3g3/8 e5/16 d5/16 c5/8 h4/8 a3c4e4/6 h4/16 h4/4 g3h3/4 a4/8 c5/8 f3a4d5/6 e5/16 d5/4 h3g3/8 e5/16 d5/16 c5/8 h4/8 c3g3e4h4/6 g5/16 g5/4 " +
		"e3/8 h3/8 g4a5/8 g5/8 d3f3c4e5/6 a4/16 e5/4 h3g3/8 d5/16 c5/16 h4/8 g4/8 a3c4e4/8 c4/16 h4/16 e4h4/8 g4/8 " +
		"h3g3/4 a4/8 c5/8 f3a4d5/8 d4/16 e5/16 a4d5/8 f3c5/8 g3h4/8 e5/16 d5/16 h4/8 g4/8 a3c4e4a4/8 c4/8 e4/8 d4/8 " +
		"a3e4c5a5/8", 40)
	if err := s.Play(&kb); err != nil {
		panic(err)
	}
}
