package windows

import (
	"github.com/kmchan2018/mindctrl/client"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/errors"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/options"
	"github.com/spf13/cobra"
)

var (
	FocusCommand = &cobra.Command{
		Use:   "focus window",
		Short: "Focus the target window",
		Long:  "Focus the target window",

		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
	}
)

func init() {
	FocusCommand.Args = func(cmd *cobra.Command, args []string) error {
		length := len(args)

		if length > 1 {
			return errors.NewExcessArgumentError()
		} else if length < 1 {
			return errors.NewMissingArgumentError("window")
		} else if options.IsId(args[0]) == false {
			return errors.NewInvalidArgumentError("window", "argument should be a valid window id")
		} else {
			return nil
		}
	}

	FocusCommand.RunE = func(cmd *cobra.Command, args []string) error {
		if transport, err := options.GetTransport(cmd); err != nil {
			return errors.WrapExecutionError(err, "cannot connect to browser")
		} else {
			window := options.ParseId(args[0])
			operation := mindctrl.FocusWindow(window)
			stdout := cmd.OutOrStdout()

			if data, err := operation.Execute(transport); err != nil {
				return errors.WrapExecutionError(err, "cannot focus window %d", window)
			} else {
				printOperationResult(stdout, window, "focused", data)
				return nil
			}
		}
	}
}
