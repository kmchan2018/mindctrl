package tabs

import (
	"fmt"
	"github.com/kmchan2018/mindctrl/client"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/errors"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/options"
	"github.com/spf13/cobra"
)

var (
	ListCommand = &cobra.Command{
		Use:   "list",
		Short: "List tabs",
		Long:  "List tabs",

		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
	}
)

func init() {
	const STATUS = "status"
	const URL = "url"
	const WINDOW = "window"

	flags := ListCommand.Flags()
	flags.String(STATUS, "", "include only tabs that have the given status")
	flags.String(URL, "", "include only tabs whose URL matches the given match pattern")
	flags.Int(WINDOW, 0, "include only tabs that appears in the given window")

	ListCommand.Args = func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return errors.NewExcessArgumentError()
		} else {
			return nil
		}
	}

	ListCommand.RunE = func(cmd *cobra.Command, args []string) error {
		if transport, err := options.GetTransport(cmd); err != nil {
			return errors.WrapExecutionError(err, "cannot connect to browser")
		} else {
			operation := mindctrl.FindTabs()
			stdout := cmd.OutOrStdout()
			flags := cmd.Flags()

			if flags.Changed(STATUS) {
				status, _ := flags.GetString(STATUS)
				operation.SetStatus(true, status)
			}

			if flags.Changed(URL) {
				url, _ := flags.GetString(URL)
				operation.SetUrl(true, url)
			}

			if flags.Changed(WINDOW) {
				window, _ := flags.GetInt(WINDOW)
				operation.SetWindowId(true, window)
			}

			if list, err := operation.Execute(transport); err != nil {
				return errors.WrapExecutionError(err, "cannot list tabs")
			} else if count := len(list); count == 0 {
				fmt.Fprintf(stdout, "No tabs found.\n\n")
				return nil
			} else {
				fmt.Fprintf(stdout, "%d tabs found:\n\n", count)

				for _, data := range list {
					fmt.Fprintf(stdout, "%12d | %s\n", data.Id, data.Url)
				}

				fmt.Fprintf(stdout, "\n")
				return nil
			}
		}
	}
}
