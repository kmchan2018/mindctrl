package windows

import (
	"github.com/kmchan2018/mindctrl/client"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/errors"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/options"
	"github.com/spf13/cobra"
)

var (
	MoveCommand = &cobra.Command{
		Use:   "move [ window ] left top",
		Short: "Move the target window",
		Long:  "Move the target window",

		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
	}
)

func init() {
	MoveCommand.Args = func(cmd *cobra.Command, args []string) error {
		length := len(args)

		if length > 3 {
			return errors.NewExcessArgumentError()
		} else if length < 1 {
			return errors.NewMissingArgumentError("left")
		} else if length < 2 {
			return errors.NewMissingArgumentError("top")
		} else if length == 3 && options.IsId(args[0]) == false {
			return errors.NewInvalidArgumentError("window", "argument should be a valid window id")
		} else if length == 2 && options.IsNumber(args[0]) == false {
			return errors.NewInvalidArgumentError("left", "argument should be a valid number")
		} else if length == 3 && options.IsNumber(args[1]) == false {
			return errors.NewInvalidArgumentError("left", "argument should be a valid number")
		} else if length == 2 && options.IsNumber(args[1]) == false {
			return errors.NewInvalidArgumentError("top", "argument should be a valid number")
		} else if length == 3 && options.IsNumber(args[2]) == false {
			return errors.NewInvalidArgumentError("top", "argument should be a valid number")
		} else {
			return nil
		}
	}

	MoveCommand.RunE = func(cmd *cobra.Command, args []string) error {
		if transport, err := options.GetTransport(cmd); err != nil {
			return errors.WrapExecutionError(err, "cannot connect to browser")
		} else {
			window := 0
			left := 0
			top := 0
			operation := mindctrl.MoveWindow(window, left, top)
			stdout := cmd.OutOrStdout()

			if len(args) == 3 {
				window = options.ParseId(args[0])
				left = options.ParseNumber(args[1])
				top = options.ParseNumber(args[2])
				operation.SetWindowId(window)
				operation.SetLeft(left)
				operation.SetTop(top)
			} else {
				if data, err := mindctrl.GetCurrentWindow().Execute(transport); err != nil {
					return errors.WrapExecutionError(err, "cannot identify current window")
				} else {
					window = data.Id
					left = options.ParseNumber(args[0])
					top = options.ParseNumber(args[1])
					operation.SetWindowId(window)
					operation.SetLeft(left)
					operation.SetTop(top)
				}
			}

			if data, err := operation.Execute(transport); err != nil {
				return errors.WrapExecutionError(err, "cannot move window %d", window)
			} else {
				printOperationResult(stdout, window, "moved", data)
				return nil
			}
		}
	}
}
