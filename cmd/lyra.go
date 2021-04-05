package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "lyra",
		Short: "Play scores on the Windsong Lyre",
		Long: `Lyra is a simple application that plays custom scors on the Windsong Lyre - toy from limited Genshin Impact event Windblume Festival.
The application sends mapped key events that simulate pressing lyre notes.`,
	}

	debug bool
)

func init() {
	rootCmd.AddCommand(EncodeCmd(), playCmd, ListenCmd())
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Print debug information.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
