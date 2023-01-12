package tabs

import (
	"github.com/kmchan2018/mindctrl/client"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/errors"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/options"
	"github.com/spf13/cobra"
)

var (
	ActivateCommand = &cobra.Command{
		Use:   "activate tab",
		Short: "Activate the target tab",
		Long:  "Activate the target tab",

		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
	}
)

func init() {
	ActivateCommand.Args = func(cmd *cobra.Command, args []string) error {
		length := len(args)

		if length > 1 {
			return errors.NewExcessArgumentError()
		} else if length < 1 {
			return errors.NewMissingArgumentError("tab")
		} else if options.IsId(args[0]) == false {
			return errors.NewInvalidArgumentError("tab", "argument should be a valid tab id")
		} else {
			return nil
		}
	}

	ActivateCommand.RunE = func(cmd *cobra.Command, args []string) error {
		if transport, err := options.GetTransport(cmd); err != nil {
			return errors.WrapExecutionError(err, "cannot connect to browser")
		} else {
			tab := options.ParseId(args[0])
			operation := mindctrl.ActivateTab(tab)
			stdout := cmd.OutOrStdout()

			if data, err := operation.Execute(transport); err != nil {
				return errors.WrapExecutionError(err, "cannot activate tab %d", tab)
			} else {
				printOperationResult(stdout, tab, "activated", data)
				return nil
			}
		}
	}
}
