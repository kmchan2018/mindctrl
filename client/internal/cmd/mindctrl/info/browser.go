package info

import (
	"fmt"
	"github.com/kmchan2018/mindctrl/client"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/errors"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/options"
	"github.com/spf13/cobra"
)

var (
	BrowserCommand = &cobra.Command{
		Use:   "browser",
		Short: "Print information about the browser",
		Long:  "Print information about the browser",

		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
	}
)

func init() {
	BrowserCommand.Args = func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return errors.NewExcessArgumentError()
		} else {
			return nil
		}
	}

	BrowserCommand.RunE = func(cmd *cobra.Command, args []string) error {
		if transport, err := options.GetTransport(cmd); err != nil {
			return errors.WrapExecutionError(err, "cannot connect to browser")
		} else if browser, err := mindctrl.GetBrowserInfo().Execute(transport); err != nil {
			return errors.WrapExecutionError(err, "cannot fetch information on the browser")
		} else {
			stdout := cmd.OutOrStdout()
			fmt.Fprintf(stdout, "Browser Name: %s\n", browser.Name)
			fmt.Fprintf(stdout, "Browser Version: %s\n", browser.Version)
			fmt.Fprintf(stdout, "\n")
			return nil
		}
	}
}
