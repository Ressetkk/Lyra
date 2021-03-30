package player

import (
	"fmt"
	"github.com/Ressetkk/windblume-lyre-player/score"
	"github.com/micmonay/keybd_event"
	"runtime"
	"time"
)

type Player struct {
	kb *keybd_event.KeyBonding
}

func New() (*Player, error) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}
	var p Player
	p.kb = &kb
	return &p, nil
}

func (p Player) Play(s *score.Score) {
	fmt.Println("Waiting 10 seconds before playing a score. Open up your game window!")
	<-time.After(10 * time.Second)
}
