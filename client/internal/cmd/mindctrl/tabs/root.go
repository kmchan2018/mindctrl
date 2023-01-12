package tabs

import (
	"github.com/spf13/cobra"
)

var (
	RootCommand *cobra.Command
)

func init() {
	RootCommand = &cobra.Command{
		Use:     "tabs",
		Aliases: []string{"tab"},
		Short:   "Manage tabs in the browser",
		Long:    "Manage tabs in the browser",
	}

	RootCommand.AddCommand(ListCommand)
	RootCommand.AddCommand(CreateCommand)
	RootCommand.AddCommand(RetrieveCommand)
	RootCommand.AddCommand(LoadCommand)
	RootCommand.AddCommand(ReloadCommand)
	RootCommand.AddCommand(ActivateCommand)
	RootCommand.AddCommand(MuteCommand)
	RootCommand.AddCommand(UnmuteCommand)
	RootCommand.AddCommand(PinCommand)
	RootCommand.AddCommand(UnpinCommand)
	RootCommand.AddCommand(MoveCommand)
	RootCommand.AddCommand(DiscardCommand)
	RootCommand.AddCommand(RemoveCommand)
}
