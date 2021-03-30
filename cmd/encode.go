package cmd

import (
	"fmt"
	"github.com/Ressetkk/lyra/player"
	"github.com/Ressetkk/lyra/score"
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
			sc, err := score.Parse(o.notes, o.name, o.author, o.tempo)
			if err != nil {
				log.Fatal(err)
			}

			initMsg := "Encoding song"
			if sc.Name != "" {
				initMsg += fmt.Sprintf(" \"%s\"", sc.Name)
			}
			if sc.Author != "" {
				initMsg += fmt.Sprintf(", by %s", sc.Author)
			}
			initMsg += "..."
			fmt.Println(initMsg)

			enc, err := score.Encode(sc)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(enc)
			if o.play {
				if p, err := player.New(debug); err != nil {
					log.Fatal(err)
				} else {
					if err := p.Play(sc); err != nil {
						log.Fatal(err)
					}
				}
			}
		},
	}

	encodeCmd.Flags().StringVar(&o.author, "author", "", "Score author")
	encodeCmd.Flags().StringVar(&o.name, "name", "", "Name of the score")
	encodeCmd.Flags().StringVarP(&o.notes, "notes", "n", "", "Notes to encode")
	encodeCmd.Flags().IntVarP(&o.tempo, "tempo", "t", 60, "Tempo of the score")
	encodeCmd.Flags().BoolVar(&o.play, "play", false, "Play the score after encoding")
	encodeCmd.MarkFlagRequired("notes")
	encodeCmd.MarkFlagRequired("tempo")

	return encodeCmd
}
