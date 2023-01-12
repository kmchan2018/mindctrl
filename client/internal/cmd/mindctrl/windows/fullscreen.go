package windows

import (
	"github.com/kmchan2018/mindctrl/client"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/errors"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/options"
	"github.com/spf13/cobra"
)

var (
	FullscreenCommand = &cobra.Command{
		Use:   "fullscreen [ window ]",
		Short: "Fullscreen the target window",
		Long:  "Fullscreen the target window",

		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
	}
)

func init() {
	FullscreenCommand.Args = func(cmd *cobra.Command, args []string) error {
		length := len(args)

		if length > 1 {
			return errors.NewExcessArgumentError()
		} else if length == 1 && options.IsId(args[0]) == false {
			return errors.NewInvalidArgumentError("window", "argument should be a valid window id")
		} else {
			return nil
		}
	}

	FullscreenCommand.RunE = func(cmd *cobra.Command, args []string) error {
		if transport, err := options.GetTransport(cmd); err != nil {
			return errors.WrapExecutionError(err, "cannot connect to browser")
		} else {
			window := 0
			operation := mindctrl.FullscreenWindow(window)
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

			if data, err := mindctrl.FullscreenWindow(window).Execute(transport); err != nil {
				return errors.WrapExecutionError(err, "cannot fullscreen window %d", window)
			} else {
				printOperationResult(stdout, window, "fullscreened", data)
				return nil
			}
		}
	}
}
