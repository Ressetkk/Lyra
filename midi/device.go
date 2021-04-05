package midi

import (
	"github.com/micmonay/keybd_event"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/reader"
	driver "gitlab.com/gomidi/rtmididrv"
	"os"
	"os/signal"
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

type Device struct {
	kb  keybd_event.KeyBonding
	drv midi.Driver
}

func NewDevice() (*Device, error) {
	dev := new(Device)
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		return nil, err
	}

	drv, err := driver.New()
	if err != nil {
		return nil, err
	}
	dev.drv = drv
	dev.kb = kb
	return dev, nil
}

func (d Device) ListenTo(devNum int) error {
	ins, err := d.drv.Ins()
	if err != nil {
		return err
	}

	in := ins[devNum]
	err = in.Open()
	if err != nil {
		return err
	}
	rd := reader.New(reader.NoLogger(), reader.NoteOn(func(p *reader.Position, channel, key, velocity uint8) {
		d.kb.SetKeys(keymap[key])
		d.kb.Launching()
	}))
	rd.ListenTo(in)

	int := make(chan os.Signal)
	signal.Notify(int, os.Interrupt)
	<-int
	return nil
}

func (d Device) Ins() ([]midi.In, error) {
	return d.drv.Ins()
}
