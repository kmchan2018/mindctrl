package downloads

import (
	"github.com/kmchan2018/mindctrl/client"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/errors"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/options"
	"github.com/spf13/cobra"
)

var (
	ResumeCommand = &cobra.Command{
		Use:   "resume download",
		Short: "Resume a download",
		Long:  "Resume a download",

		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
	}
)

func init() {
	ResumeCommand.Args = func(cmd *cobra.Command, args []string) error {
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

	ResumeCommand.RunE = func(cmd *cobra.Command, args []string) error {
		if transport, err := options.GetTransport(cmd); err != nil {
			return errors.WrapExecutionError(err, "cannot connect to browser")
		} else {
			download := options.ParseId(args[0])
			operation := mindctrl.ResumeDownload(download)
			stdout := cmd.OutOrStdout()

			if data, err := operation.Execute(transport); err != nil {
				return errors.WrapExecutionError(err, "cannot resume download %d", download)
			} else {
				printOperationResult(stdout, download, "resumed", data)
				return nil
			}
		}
	}
}
