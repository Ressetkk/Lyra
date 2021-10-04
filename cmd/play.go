package cmd

import (
	"errors"
	"fmt"
	"github.com/Ressetkk/lyra/pkg/lyre"
	"github.com/spf13/cobra"
	"gitlab.com/gomidi/midi/player"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play the MIDI file.",
	Long: `Play the given MIDI file.
Source can be either http URL or local file.
After running the program you'll have 10 seconds before it starts generating key events.`,
	Run: play,
}

func play(_ *cobra.Command, args []string) {
	time.Sleep(5 * time.Second)
	file := args[0]
	stop := make(chan bool)
	finished := make(chan bool)

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)

	l, err := lyre.New()
	if err != nil {
		log.Fatalln(err)
	}
	f, err := openStream(file)
	if err != nil {
		log.Fatalln(err)
	}

	p, err := player.New(f)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Playing file", file)
	p.PlayAllTo(l, stop, finished)

	select {
	case <-finished:
		fmt.Println("Playback finished.")
		break
	case <-sig:
		fmt.Println("Interrupting playback.")
		stop <- true
		break
	}
	fmt.Println("Exiting.")
}

func openStream(file string) (io.ReadCloser, error) {
	if len(file) == 0 {
		return nil, errors.New("file name cannot be empty")
	}
	if strings.HasPrefix(file, "https://") || strings.HasPrefix(file, "http://") {
		// we treat it as HTTP stream, so we open it as http stream
		resp, err := http.Get(file)
		if err != nil {
			return nil, err
		}
		return resp.Body, nil
	}

	if _, err := os.Stat(file); err != nil {
		return nil, err
	} else {
		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		return f, nil
	}
}
