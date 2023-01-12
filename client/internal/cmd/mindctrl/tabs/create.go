package tabs

import (
	"github.com/kmchan2018/mindctrl/client"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/errors"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/options"
	"github.com/spf13/cobra"
)

var (
	CreateCommand = &cobra.Command{
		Use:     "create",
		Aliases: []string{"open"},
		Short:   "Create a new tab",
		Long:    "Create a new tab",

		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
	}
)

func init() {
	const ACTIVE = "active"
	const URL = "url"
	const WINDOW = "window"

	flags := CreateCommand.Flags()
	flags.Bool(ACTIVE, false, "whether the tab is activated upon creation")
	flags.String(URL, "", "url to be loaded in the new tab")
	flags.Int(WINDOW, 0, "window where the new window is created in")

	CreateCommand.Args = func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return errors.NewExcessArgumentError()
		} else {
			return nil
		}
	}

	CreateCommand.RunE = func(cmd *cobra.Command, args []string) error {
		if transport, err := options.GetTransport(cmd); err != nil {
			return errors.WrapExecutionError(err, "cannot connect to browser")
		} else {
			operation := mindctrl.CreateTab()
			stdout := cmd.OutOrStdout()
			flags := cmd.Flags()

			if flags.Changed(ACTIVE) {
				active, _ := flags.GetBool(ACTIVE)
				operation.SetActive(true, active)
			}

			if flags.Changed(URL) {
				url, _ := flags.GetString(URL)
				operation.SetUrl(true, url)
			}

			if flags.Changed(WINDOW) {
				window, _ := flags.GetInt(WINDOW)
				operation.SetWindowId(true, window)
			}

			if data, err := operation.Execute(transport); err != nil {
				return errors.WrapExecutionError(err, "cannot create new tab")
			} else {
				tab := data.Id
				printOperationResult(stdout, tab, "created", data)
				return nil
			}
		}
	}
}
