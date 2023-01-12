package tabs

import (
	"github.com/kmchan2018/mindctrl/client"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/errors"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/options"
	"github.com/spf13/cobra"
)

var (
	MoveCommand = &cobra.Command{
		Use:   "move [ tab ] index",
		Short: "Move the target tab",
		Long:  "Move the target tab",

		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
	}
)

func init() {
	MoveCommand.Args = func(cmd *cobra.Command, args []string) error {
		length := len(args)

		if length > 2 {
			return errors.NewExcessArgumentError()
		} else if length < 1 {
			return errors.NewMissingArgumentError("index")
		} else if length == 2 && options.IsId(args[0]) == false {
			return errors.NewInvalidArgumentError("tab", "argument should be a valid tab id")
		} else if length == 1 && options.IsNumber(args[0]) == false {
			return errors.NewInvalidArgumentError("index", "argument should be a valid number")
		} else if length == 2 && options.IsNumber(args[1]) == false {
			return errors.NewInvalidArgumentError("index", "argument should be a valid number")
		} else {
			return nil
		}
	}

	MoveCommand.RunE = func(cmd *cobra.Command, args []string) error {
		if transport, err := options.GetTransport(cmd); err != nil {
			return errors.WrapExecutionError(err, "cannot connect to browser")
		} else {
			tab := 0
			index := 0
			operation := mindctrl.MoveTab(tab, index)
			stdout := cmd.OutOrStdout()

			if len(args) == 2 {
				tab = options.ParseId(args[0])
				index = options.ParseNumber(args[1])
				operation.SetTabId(tab)
				operation.SetIndex(index)
			} else {
				if data, err := mindctrl.GetCurrentTab().Execute(transport); err != nil {
					return errors.WrapExecutionError(err, "cannot identify current tab")
				} else {
					tab = data.Id
					index = options.ParseNumber(args[0])
					operation.SetTabId(tab)
					operation.SetIndex(index)
				}
			}

			if data, err := operation.Execute(transport); err != nil {
				return errors.WrapExecutionError(err, "cannot move tab %d", tab)
			} else {
				printOperationResult(stdout, tab, "moved", data)
				return nil
			}
		}
	}
}
