package lyre

import (
	"github.com/micmonay/keybd_event"
	"github.com/pkg/errors"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/midimessage/channel"
	"runtime"
	"time"
)

var keymap = map[uint8]int{
	48: keybd_event.VK_Z,
	50: keybd_event.VK_X,
	52: keybd_event.VK_C,
	53: keybd_event.VK_V,
	55: keybd_event.VK_B,
	57: keybd_event.VK_N,
	59: keybd_event.VK_M,
	60: keybd_event.VK_A,
	62: keybd_event.VK_S,
	64: keybd_event.VK_D,
	65: keybd_event.VK_F,
	67: keybd_event.VK_G,
	69: keybd_event.VK_H,
	71: keybd_event.VK_J,
	72: keybd_event.VK_Q,
	74: keybd_event.VK_W,
	76: keybd_event.VK_E,
	77: keybd_event.VK_R,
	79: keybd_event.VK_T,
	81: keybd_event.VK_Y,
	83: keybd_event.VK_U,
}

type Lyre struct {
	kb keybd_event.KeyBonding
}

func New() (*Lyre, error) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		return nil, err
	}
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}
	l := Lyre{
		kb: kb,
	}
	return &l, nil
}

func (l Lyre) Write(m midi.Message) error {
	switch v := m.(type) {
	case channel.NoteOn:
		return l.press(v.Key())
	default:
		return nil
	}
}

func (l Lyre) press(key uint8) error {
	k, ok := keymap[key]
	if !ok {
		return nil // unallowed key - do nothing
	}
	l.kb.SetKeys(k)
	return errors.Wrap(l.kb.Launching(), "press key")
}
