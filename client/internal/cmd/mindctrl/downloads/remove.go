package downloads

import (
	"github.com/kmchan2018/mindctrl/client"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/errors"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/options"
	"github.com/spf13/cobra"
)

var (
	RemoveCommand = &cobra.Command{
		Use:   "remove download",
		Short: "Remove a download",
		Long:  "Remove a download",

		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
	}
)

func init() {
	RemoveCommand.Args = func(cmd *cobra.Command, args []string) error {
		length := len(args)

		if length > 1 {
			return errors.NewExcessArgumentError()
		} else if length < 1 {
			return errors.NewMissingArgumentError("download")
		} else if options.IsId(args[0]) == false {
			return errors.NewInvalidArgumentError("download", "argument should be a valid download id")
		} else {
			return nil
		}
	}

	RemoveCommand.RunE = func(cmd *cobra.Command, args []string) error {
		if transport, err := options.GetTransport(cmd); err != nil {
			return errors.WrapExecutionError(err, "cannot connect to browser")
		} else {
			download := options.ParseId(args[0])
			operation := mindctrl.RemoveDownload(download)
			stdout := cmd.OutOrStdout()

			if err := operation.Execute(transport); err != nil {
				return errors.WrapExecutionError(err, "cannot remove download %d", download)
			} else {
				printOperationResult(stdout, download, "removed", nil)
				return nil
			}
		}
	}
}
