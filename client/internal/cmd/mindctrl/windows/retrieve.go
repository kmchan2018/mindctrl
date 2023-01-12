package windows

import (
	"github.com/kmchan2018/mindctrl/client"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/errors"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/options"
	"github.com/spf13/cobra"
)

var (
	RetrieveCommand = &cobra.Command{
		Use:     "retrieve [ window ]",
		Aliases: []string{"get", "show"},
		Short:   "Retrieve the target window",
		Long:    "Retrieve the target window",

		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
	}
)

func init() {
	RetrieveCommand.Args = func(cmd *cobra.Command, args []string) error {
		length := len(args)

		if length > 1 {
			return errors.NewArgumentError("Excess argument found")
		} else if length == 1 && options.IsId(args[0]) == false {
			return errors.NewArgumentError("Argument 'window' should be a valid window id")
		} else {
			return nil
		}
	}

	RetrieveCommand.RunE = func(cmd *cobra.Command, args []string) error {
		if transport, err := options.GetTransport(cmd); err != nil {
			return errors.WrapExecutionError(err, "Cannot connect to browser")
		} else {
			window := 0
			operation := mindctrl.GetWindow(window)
			stdout := cmd.OutOrStdout()

			if len(args) > 0 {
				window = options.ParseId(args[0])
				operation.SetWindowId(window)
			} else {
				if data, err := mindctrl.GetCurrentWindow().Execute(transport); err != nil {
					return errors.WrapExecutionError(err, "Cannot identify current window")
				} else {
					window = data.Id
					operation.SetWindowId(window)
				}
			}

			if data, err := operation.Execute(transport); err != nil {
				return errors.WrapExecutionError(err, "Cannot retrieve window %d", window)
			} else {
				printOperationResult(stdout, window, "retrieved", data)
				return nil
			}
		}
	}
}
