package cable

import (
	"fmt"
	"github.com/Ressetkk/lyra/pkg/lyre"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/reader"
	driver "gitlab.com/gomidi/rtmididrv"
	"os"
	"os/signal"
)

type Cable struct {
	drv midi.Driver
}

func New() (*Cable, error) {
	dev := new(Cable)
	drv, err := driver.New()
	if err != nil {
		return nil, err
	}
	dev.drv = drv
	return dev, nil
}

func (d Cable) Bridge(devNum int) error {
	_, err := lyre.New()
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
		fmt.Println(key)
	}))

	rd.ListenTo(in)
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
	<-sig
	return nil
}

func (d Cable) Ins() ([]midi.In, error) {
	return d.drv.Ins()
}
