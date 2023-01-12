package downloads

import (
	"github.com/kmchan2018/mindctrl/client"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/errors"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/options"
	"github.com/spf13/cobra"
)

var (
	CreateCommand = &cobra.Command{
		Use:   "create url filename",
		Short: "Create a new download",
		Long:  "Create a new download",

		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
	}
)

func init() {
	REFERRER := "referrer"

	flags := CreateCommand.Flags()
	flags.String(REFERRER, "", "referrer used in download requests")

	CreateCommand.Args = func(cmd *cobra.Command, args []string) error {
		length := len(args)

		if length > 2 {
			return errors.NewExcessArgumentError()
		} else if length < 1 {
			return errors.NewMissingArgumentError("url")
		} else if length < 2 {
			return errors.NewMissingArgumentError("filename")
		} else if args[0] == "" {
			return errors.NewInvalidArgumentError("url", "argument should be a valid url")
		} else if args[1] == "" {
			return errors.NewInvalidArgumentError("url", "argument should be a valid file path")
		} else {
			return nil
		}
	}

	CreateCommand.RunE = func(cmd *cobra.Command, args []string) error {
		if transport, err := options.GetTransport(cmd); err != nil {
			return errors.WrapExecutionError(err, "cannot connect to browser")
		} else {
			url := args[0]
			filename := args[1]
			operation := mindctrl.CreateDownload(url, filename)
			stdout := cmd.OutOrStdout()
			flags := cmd.Flags()

			if flags.Changed(REFERRER) {
				referrer, _ := flags.GetString(REFERRER)
				operation.SetReferrer(true, referrer)
			}

			if data, err := operation.Execute(transport); err != nil {
				return errors.WrapExecutionError(err, "cannot create new download of url %s to file %s", url, filename)
			} else {
				download := data.Id
				printOperationResult(stdout, download, "created", data)
				return nil
			}
		}
	}
}
