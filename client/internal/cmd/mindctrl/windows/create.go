package windows

import (
	"github.com/kmchan2018/mindctrl/client"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/errors"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/options"
	"github.com/spf13/cobra"
)

var (
	CreateCommand = &cobra.Command{
		Use:     "create",
		Aliases: []string{"open"},
		Short:   "Create a new window",
		Long:    "Create a new window",

		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
	}
)

func init() {
	CreateCommand.Args = func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return errors.NewExcessArgumentError()
		} else {
			return nil
		}
	}

	CreateCommand.RunE = func(cmd *cobra.Command, args []string) error {
		if transport, err := options.GetTransport(cmd); err != nil {
			return errors.WrapExecutionError(err, "cannot connect to browser")
		} else {
			operation := mindctrl.CreateWindow()
			stdout := cmd.OutOrStdout()

			if data, err := operation.Execute(transport); err != nil {
				return errors.WrapExecutionError(err, "cannot create new window")
			} else {
				window := data.Id
				printOperationResult(stdout, window, "created", data)
				return nil
			}
		}
	}
}
