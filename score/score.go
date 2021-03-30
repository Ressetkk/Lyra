package score

import (
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
	Tempo  int    `json:"tempo,omitempty"`
	Notes  []Note `json:"notes,omitempty"`
}

type Note struct {
	Beat  int    `json:"beat,omitempty"`
	Notes []int  `json:"notes,omitempty"`
	Mode  string `json:"mode,omitempty"`
}

type DecodeError struct {
	err error
}

func (de DecodeError) Error() string {
	return fmt.Sprintf("score decode error: %v", de.err)
}

func Parse(exp, name, author string, tempo int) (*Score, error) {
	var score Score
	score.Tempo = tempo
	score.Name = name
	score.Author = author

	if n, err := ParseNotes(exp); err != nil {
		return nil, err
	} else {
		score.Notes = n
	}

	return &score, nil
}

func ParseNotes(exp string) ([]Note, error) {
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
	mode, s, err := splitMode(s)
	if err != nil {
		return nil, err
	}
	n.Mode = mode

	sn := strings.Split(s, "/")
	if len(sn) != 2 {
		return nil, fmt.Errorf("couldn't split note: %v", s)
	}
	if v, err := parse(sn[0]); err != nil {
		return nil, err
	} else {
		n.Notes = v
	}
	// TODO convert to float to easy allow extended notes
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

func splitMode(s string) (string, string, error) {
	if strings.Contains(s, ":") {
		sn := strings.Split(s, ":")
		if len(sn) != 2 {
			return "", "", fmt.Errorf("couldn't extract mode: %v", s)
		}
		return sn[0], sn[1], nil
	}
	return "", s, nil
}
