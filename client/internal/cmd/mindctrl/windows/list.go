package windows

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
		Short: "List all windows",
		Long:  "List all windows in the browser",

		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
	}
)

func init() {
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
			operation := mindctrl.FindWindows()
			stdout := cmd.OutOrStdout()

			if list, err := operation.Execute(transport); err != nil {
				return errors.WrapExecutionError(err, "cannot list windows")
			} else if count := len(list); count == 0 {
				fmt.Fprintf(stdout, "No windows found.\n\n")
				return nil
			} else {
				fmt.Fprintf(stdout, "%d windows found:\n\n", count)

				for _, data := range list {
					id := data.Id
					title := "(Empty)"

					if len(data.Tabs) > 0 {
						title = data.Tabs[0].Title
					}

					for _, data2 := range data.Tabs {
						if data2.Active {
							title = data2.Title
						}
					}

					fmt.Fprintf(stdout, "%12d | %s\n", id, title)
				}

				fmt.Fprintf(stdout, "\n")
				return nil
			}
		}
	}
}
