package score

import (
	"fmt"
	"github.com/micmonay/keybd_event"
	"log"
	"strconv"
	"strings"
	"time"
)

var keymap = map[string]int{
	"c3": keybd_event.VK_Z,
	"d3": keybd_event.VK_X,
	"e3": keybd_event.VK_C,
	"f3": keybd_event.VK_V,
	"g3": keybd_event.VK_B,
	"a3": keybd_event.VK_N,
	"h3": keybd_event.VK_M,
	"c4": keybd_event.VK_A,
	"d4": keybd_event.VK_S,
	"e4": keybd_event.VK_D,
	"f4": keybd_event.VK_F,
	"g4": keybd_event.VK_G,
	"a4": keybd_event.VK_H,
	"h4": keybd_event.VK_J,
	"c5": keybd_event.VK_Q,
	"d5": keybd_event.VK_W,
	"e5": keybd_event.VK_E,
	"f5": keybd_event.VK_R,
	"g5": keybd_event.VK_T,
	"a5": keybd_event.VK_Y,
	"h5": keybd_event.VK_U,
}

type Score struct {
	Exp string
	Tempo int
}

func NewScore(exp string, tempo int) *Score {
	return &Score{Exp: exp, Tempo: tempo}
}

type Note struct {
	Beat int
	Note []int
}

// TODO move it to separate player package
func (s Score) Play(kb *keybd_event.KeyBonding) error {
	scoreList := strings.Split(s.Exp, " ")
	bpm := time.Minute / time.Duration(s.Tempo)
	log.Printf("starting play bpm=%v", bpm)
	for _, note := range scoreList {
		pNote, err := ParseNote(note)
		if err != nil {
			return err
		}
		kb.SetKeys(pNote.Note...)
		kb.Launching()

		sl := bpm / time.Duration(pNote.Beat) * 4
		log.Printf("note=%v beat=%v sleep=%v", pNote.Note, pNote.Beat, sl)
		time.Sleep(sl)
	}

	return nil
}

func ParseNote(s string) (*Note, error) {
	var n Note
	sn:= strings.Split(s, "/")
	if len(sn) != 2 {
		return nil, fmt.Errorf("couldn't split note: %v", s)
	}
	if v, err := parse(sn[0]); err != nil {
		return nil, err
	} else {
		n.Note = v
	}

	pb, err := strconv.Atoi(sn[1])
	if err != nil {
		return nil, err
	}
	n.Beat = pb

	return &n, nil
}

func parse(s string) ([]int, error) {
	var elems []int
	for s != "" {
		n := s[:2]
		s = s[2:]

		if v, ok := keymap[n]; !ok {
			return nil, fmt.Errorf("unknown note: %v", n)
		} else {
			elems = append(elems, v)
		}
	}
	return elems, nil
}
