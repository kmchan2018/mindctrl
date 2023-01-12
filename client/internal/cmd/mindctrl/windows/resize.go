package windows

import (
	"github.com/kmchan2018/mindctrl/client"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/errors"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/options"
	"github.com/spf13/cobra"
)

var (
	ResizeCommand = &cobra.Command{
		Use:   "resize [ window ] width height",
		Short: "Resize the target window",
		Long:  "Resize the target window",

		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
	}
)

func init() {
	ResizeCommand.Args = func(cmd *cobra.Command, args []string) error {
		length := len(args)

		if length > 3 {
			return errors.NewExcessArgumentError()
		} else if length < 1 {
			return errors.NewMissingArgumentError("width")
		} else if length < 2 {
			return errors.NewMissingArgumentError("height")
		} else if length == 3 && options.IsId(args[0]) == false {
			return errors.NewInvalidArgumentError("window", "argument should be a valid window id")
		} else if length == 2 && options.IsNumber(args[0]) == false {
			return errors.NewInvalidArgumentError("width", "argument should be a valid number")
		} else if length == 3 && options.IsNumber(args[1]) == false {
			return errors.NewInvalidArgumentError("width", "argument should be a valid number")
		} else if length == 2 && options.IsNumber(args[1]) == false {
			return errors.NewInvalidArgumentError("height", "argument should be a valid number")
		} else if length == 3 && options.IsNumber(args[2]) == false {
			return errors.NewInvalidArgumentError("height", "argument should be a valid number")
		} else {
			return nil
		}
	}

	ResizeCommand.RunE = func(cmd *cobra.Command, args []string) error {
		if transport, err := options.GetTransport(cmd); err != nil {
			return errors.WrapExecutionError(err, "cannot connect to browser")
		} else {
			window := 0
			width := 0
			height := 0
			operation := mindctrl.ResizeWindow(window, width, height)
			stdout := cmd.OutOrStdout()

			if len(args) == 3 {
				window = options.ParseId(args[0])
				width = options.ParseNumber(args[1])
				height = options.ParseNumber(args[2])
				operation.SetWindowId(window)
				operation.SetWidth(width)
				operation.SetHeight(height)
			} else {
				if data, err := mindctrl.GetCurrentWindow().Execute(transport); err != nil {
					return errors.WrapExecutionError(err, "cannot identify current window")
				} else {
					window = data.Id
					width = options.ParseNumber(args[0])
					height = options.ParseNumber(args[1])
					operation.SetWindowId(window)
					operation.SetWidth(width)
					operation.SetHeight(height)
				}
			}

			if data, err := operation.Execute(transport); err != nil {
				return errors.WrapExecutionError(err, "cannot resize window %d", window)
			} else {
				printOperationResult(stdout, window, "resized", data)
				return nil
			}
		}
	}
}
