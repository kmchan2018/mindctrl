package downloads

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
		Short: "List downloads",
		Long:  "List downloads",

		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
	}
)

func init() {
	const STATE = "state"
	const URL = "url"

	flags := ListCommand.Flags()
	flags.String(STATE, "", "includes only downloads that have the given state")
	flags.String(URL, "", "includes only downloads whose URL matches the given match pattern")

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
			operation := mindctrl.FindDownloads()
			stdout := cmd.OutOrStdout()
			flags := cmd.Flags()

			if flags.Changed(STATE) {
				state, _ := flags.GetString(STATE)
				operation.SetUrl(true, state)
			}

			if flags.Changed(URL) {
				url, _ := flags.GetString(URL)
				operation.SetUrl(true, url)
			}

			if list, err := operation.Execute(transport); err != nil {
				return errors.WrapExecutionError(err, "cannot list downloads")
			} else if count := len(list); count == 0 {
				fmt.Fprintf(stdout, "No downloads found.\n\n")
				return nil
			} else {
				fmt.Fprintf(stdout, "%d downloads found:\n\n", count)

				for _, data := range list {
					fmt.Fprintf(stdout, "%12d | %-12s | %s <= %s\n", data.Id, data.State, data.Filename, data.Url)
				}

				fmt.Fprintf(stdout, "\n")
				return nil
			}
		}
	}
}
