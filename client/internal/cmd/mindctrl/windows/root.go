package windows

import (
	"github.com/spf13/cobra"
)

var (
	RootCommand *cobra.Command
)

func init() {
	RootCommand = &cobra.Command{
		Use:     "windows",
		Aliases: []string{"window"},
		Short:   "Manage windows in the browser",
		Long:    "Manage windows in the browser",
	}

	RootCommand.AddCommand(ListCommand)
	RootCommand.AddCommand(CreateCommand)
	RootCommand.AddCommand(RetrieveCommand)
	RootCommand.AddCommand(MoveCommand)
	RootCommand.AddCommand(ResizeCommand)
	RootCommand.AddCommand(MinimizeCommand)
	RootCommand.AddCommand(MaximizeCommand)
	RootCommand.AddCommand(FullscreenCommand)
	RootCommand.AddCommand(RestoreCommand)
	RootCommand.AddCommand(FocusCommand)
	RootCommand.AddCommand(RemoveCommand)
}
