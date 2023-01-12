package downloads

import (
	"github.com/spf13/cobra"
)

var (
	RootCommand *cobra.Command
)

func init() {
	RootCommand = &cobra.Command{
		Use:     "downloads",
		Aliases: []string{"download"},
		Short:   "Manage downloads in the browser",
		Long:    "Manage downloads in the browser",
	}

	RootCommand.AddCommand(ListCommand)
	RootCommand.AddCommand(CreateCommand)
	RootCommand.AddCommand(PauseCommand)
	RootCommand.AddCommand(ResumeCommand)
	RootCommand.AddCommand(CancelCommand)
	RootCommand.AddCommand(RemoveCommand)
}
