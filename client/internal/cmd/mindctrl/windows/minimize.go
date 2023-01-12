package windows

import (
	"github.com/kmchan2018/mindctrl/client"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/errors"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/options"
	"github.com/spf13/cobra"
)

var (
	MinimizeCommand = &cobra.Command{
		Use:   "minimize [ window ]",
		Short: "Minimize the target window",
		Long:  "Minimize the target window",

		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
	}
)

func init() {
	MinimizeCommand.Args = func(cmd *cobra.Command, args []string) error {
		length := len(args)

		if length > 1 {
			return errors.NewExcessArgumentError()
		} else if length == 1 && options.IsId(args[0]) == false {
			return errors.NewInvalidArgumentError("window", "argument should be a valid window id")
		} else {
			return nil
		}
	}

	MinimizeCommand.RunE = func(cmd *cobra.Command, args []string) error {
		if transport, err := options.GetTransport(cmd); err != nil {
			return errors.WrapExecutionError(err, "cannot connect to browser")
		} else {
			window := 0
			operation := mindctrl.MinimizeWindow(window)
			stdout := cmd.OutOrStdout()

			if len(args) > 0 {
				window = options.ParseId(args[0])
				operation.SetWindowId(window)
			} else {
				if data, err := mindctrl.GetCurrentWindow().Execute(transport); err != nil {
					return errors.WrapExecutionError(err, "cannot identify current window")
				} else {
					window = data.Id
					operation.SetWindowId(window)
				}
			}

			if data, err := operation.Execute(transport); err != nil {
				return errors.WrapExecutionError(err, "cannot minimize window %d", window)
			} else {
				printOperationResult(stdout, window, "minimized", data)
				return nil
			}
		}
	}
}
