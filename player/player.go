package player

import (
	"fmt"
	"github.com/Ressetkk/lyra/score"
	"github.com/micmonay/keybd_event"
	"log"
	"runtime"
	"time"
)

type Player struct {
	kb    *keybd_event.KeyBonding
	Debug bool
}

func New(debug bool) (*Player, error) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		return nil, err
	}
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}
	var p Player
	p.kb = &kb
	p.Debug = debug
	return &p, nil
}

func (p Player) Play(s *score.Score) error {
	fmt.Println("Waiting 10 seconds before playing a score. Open up your game window!")
	<-time.After(10 * time.Second)

	noteChan := make(chan score.Note)
	errChan := make(chan error)
	go func() {
		defer close(noteChan)
		bpm := time.Minute / time.Duration(s.Tempo)

		for _, n := range s.Notes {
			// TODO add possible delay between chord notes - if set play each note with 10ms delay
			noteChan <- n
			<-time.After(bpm / time.Duration(n.Beat) * 4)
		}
	}()

	for {
		select {
		case err := <-errChan:
			return err
		case n, ok := <-noteChan:
			if ok {
				if p.Debug {
					log.Printf("note=%v beat=%v mode=%v", n.Notes, n.Beat, n.Mode)
				}
				p.kb.SetKeys(n.Notes...)
				if err := p.kb.Launching(); err != nil {
					return fmt.Errorf("playback error: %w", err)
				}
			} else {
				if p.Debug {
					log.Println("Finished playing song.")
				}
				return nil
			}
		}
	}
}
