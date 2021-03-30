package score

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/micmonay/keybd_event"
	"strconv"
	"strings"
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
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
	Tempo  int    `json:"tempo"`
	Notes  string `json:"notes"`
}

type Note struct {
	Beat  int   `json:"beat"`
	Notes []int `json:"notes"`
}

type DecodeError struct {
	err error
}

func (de DecodeError) Error() string {
	return fmt.Sprintf("score decode error: %v", de.err)
}

// TODO move it to separate player package
//func (s Score) Play(kb *keybd_event.KeyBonding) error {
//	scoreList := strings.Split(s.Exp, " ")
//	bpm := time.Minute / time.Duration(s.Tempo)
//	log.Printf("starting play bpm=%v", bpm)
//	for _, note := range scoreList {
//		pNote, err := ParseNote(note)
//		if err != nil {
//			return err
//		}
//		kb.SetKeys(pNote.Notes...)
//		kb.Launching()
//
//		sl := bpm / time.Duration(pNote.Beat) * 4
//		log.Printf("note=%v beat=%v sleep=%v", pNote.Notes, pNote.Beat, sl)
//		time.Sleep(sl)
//	}
//
//	return nil
//}

func Parse(exp, name, author string, tempo int) (*Score, error) {
	var score Score
	score.Tempo = tempo
	score.Name = name
	score.Author = author

	if _, err := parseNotes(exp); err != nil {
		return nil, err
	}

	score.Notes = exp
	return &score, nil
}

func parseNotes(exp string) ([]Note, error) {
	var notes []Note
	sepNotes := strings.Split(exp, " ")
	for _, note := range sepNotes {
		pNote, err := ParseNote(note)
		if err != nil {
			return nil, fmt.Errorf("note parse error: %w", err)
		}
		notes = append(notes, *pNote)
	}
	return notes, nil
}

func ParseNote(s string) (*Note, error) {
	var n Note
	sn := strings.Split(s, "/")
	if len(sn) != 2 {
		return nil, fmt.Errorf("couldn't split note: %v", s)
	}
	if v, err := parse(sn[0]); err != nil {
		return nil, err
	} else {
		n.Notes = v
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

func Encode(s *Score) (string, error) {
	jnotes, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(jnotes), nil
}

func Decode(exp string) (*Score, error) {
	dec, err := base64.StdEncoding.DecodeString(exp)
	if err != nil {
		return nil, DecodeError{err}
	}
	var score Score
	if err := json.Unmarshal(dec, &score); err != nil {
		return nil, DecodeError{err}
	}
	return &score, nil
}
