package cmd

import (
	"fmt"
	"github.com/Ressetkk/lyra/midi"
	"github.com/spf13/cobra"
)

var (
	dev    *midi.Device
	devNum int
)

func ListenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "listen",
		Short: "Listen to MIDI controller and translate it to key presses",
		Long: `Listen to MIDI controller events to play songs.
The application will translate every key to specific key press events.

To list available MIDI inputs type "lyra listen list".`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			d, err := midi.NewDevice()
			if err != nil {
				panic(err)
			}
			dev = d
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Listening to the keypresses on the device %v...\n", devNum)
			err := dev.ListenTo(devNum)
			if err != nil {
				panic(err)
			}
		},
	}
	cmd.Flags().IntVarP(&devNum, "device-number", "d", 0, "MIDI Input device to listen to.")

	list := &cobra.Command{
		Use:   "list",
		Short: "List available MIDI devices.",
		Long:  "List available MIDI devices that you can use as keyboard for Lyre.",
		Run: func(cmd *cobra.Command, args []string) {
			ins, err := dev.Ins()
			if err != nil {
				panic(err)
			}
			fmt.Println("Available MIDI input devices:")
			for k, v := range ins {
				fmt.Printf("%v. %s\n", k, v)
			}
		},
	}
	cmd.AddCommand(list)
	return cmd
}
