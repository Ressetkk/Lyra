package cmd

import (
	"fmt"
	"github.com/Ressetkk/windblume-lyre-player/score"
	"github.com/spf13/cobra"
	"log"
)

type encodeOpts struct {
	notes  string
	tempo  int
	name   string
	author string
	play   bool
}

func EncodeCmd() *cobra.Command {
	var o encodeOpts
	encodeCmd := &cobra.Command{
		Use:   "encode",
		Short: "Encode raw score to base64 format",
		Long: `Encode will parse the raw score and tempo to the base64-encoded string which you can play and share with others.
You can optionally provide the name and the author of the score.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Encoding score: %s, by %s...\n", o.name, o.author)
			sc, err := score.Parse(o.notes, o.name, o.author, o.tempo)
			if err != nil {
				log.Fatal(err)
			}
			enc, err := score.Encode(sc)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(enc)
			if o.play {
				// TODO play a song
			}
		},
	}

	encodeCmd.Flags().StringVar(&o.author, "author", "", "Score author")
	encodeCmd.Flags().StringVar(&o.name, "name", "", "Name of the score")
	encodeCmd.Flags().StringVarP(&o.notes, "notes", "n", "", "Notes to encode")
	encodeCmd.Flags().IntVarP(&o.tempo, "tempo", "t", 60, "Tempo of the score")
	encodeCmd.Flags().BoolVar(&o.play, "play", false, "Play the score after encoding")
	encodeCmd.MarkFlagRequired("name")
	encodeCmd.MarkFlagRequired("tempo")

	return encodeCmd
}
