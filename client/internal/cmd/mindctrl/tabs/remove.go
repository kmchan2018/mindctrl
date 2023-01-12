package tabs

import (
	"github.com/kmchan2018/mindctrl/client"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/errors"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/options"
	"github.com/spf13/cobra"
)

var (
	RemoveCommand = &cobra.Command{
		Use:     "remove [ tab ]",
		Aliases: []string{"close"},
		Short:   "Remove the target tab",
		Long:    "Remove the target tab",

		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
	}
)

func init() {
	RemoveCommand.Args = func(cmd *cobra.Command, args []string) error {
		length := len(args)

		if length > 1 {
			return errors.NewExcessArgumentError()
		} else if length == 1 && options.IsId(args[0]) == false {
			return errors.NewInvalidArgumentError("tab", "argument should be a valid tab id")
		} else {
			return nil
		}
	}

	RemoveCommand.RunE = func(cmd *cobra.Command, args []string) error {
		if transport, err := options.GetTransport(cmd); err != nil {
			return errors.WrapExecutionError(err, "cannot connect to browser")
		} else {
			tab := 0
			operation := mindctrl.RemoveTab(tab)
			stdout := cmd.OutOrStdout()

			if len(args) > 0 {
				tab = options.ParseId(args[0])
				operation.SetTabId(tab)
			} else {
				if data, err := mindctrl.GetCurrentTab().Execute(transport); err != nil {
					return errors.WrapExecutionError(err, "cannot identify current tab")
				} else {
					tab = data.Id
					operation.SetTabId(tab)
				}
			}

			if err := operation.Execute(transport); err != nil {
				return errors.WrapExecutionError(err, "cannot remove tab %d", tab)
			} else {
				printOperationResult(stdout, tab, "removed", nil)
				return nil
			}
		}
	}
}
