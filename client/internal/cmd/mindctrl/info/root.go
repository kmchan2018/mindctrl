package info

import (
	"github.com/spf13/cobra"
)

var (
	RootCommand *cobra.Command
)

func init() {
	RootCommand = &cobra.Command{
		Use:   "info",
		Short: "Print information",
		Long:  "Print information",
	}

	RootCommand.AddCommand(AllCommand)
	RootCommand.AddCommand(BrowserCommand)
	RootCommand.AddCommand(PlatformCommand)
}
