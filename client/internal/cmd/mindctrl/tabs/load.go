package tabs

import (
	"github.com/kmchan2018/mindctrl/client"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/errors"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/options"
	"github.com/spf13/cobra"
)

var (
	LoadCommand = &cobra.Command{
		Use:   "load [ tab ] url",
		Short: "Load the given URL in the target tab",
		Long:  "Load the given URL in the target tab",

		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
	}
)

func init() {
	const REPLACE = "replace"

	flags := LoadCommand.Flags()
	flags.Bool("replace", false, "replace the current page with incoming page in the history stack of the tab")

	LoadCommand.Args = func(cmd *cobra.Command, args []string) error {
		length := len(args)

		if length > 2 {
			return errors.NewExcessArgumentError()
		} else if length < 0 {
			return errors.NewMissingArgumentError("url")
		} else if length == 2 && options.IsId(args[0]) == false {
			return errors.NewInvalidArgumentError("tab", "argument should be a valid tab id")
		} else if length == 2 && args[1] == "" {
			return errors.NewInvalidArgumentError("url", "argument should be a valid url")
		} else if length == 1 && args[0] == "" {
			return errors.NewInvalidArgumentError("url", "argument should be a valid url")
		} else {
			return nil
		}
	}

	LoadCommand.RunE = func(cmd *cobra.Command, args []string) error {
		if transport, err := options.GetTransport(cmd); err != nil {
			return errors.WrapExecutionError(err, "cannot connect to browser")
		} else {
			tab := 0
			url := ""
			operation := mindctrl.LoadTab(tab, url)
			stdout := cmd.OutOrStdout()
			flags := cmd.Flags()

			if len(args) > 1 {
				tab = options.ParseId(args[0])
				url = args[1]
				operation.SetTabId(tab)
				operation.SetUrl(url)
			} else {
				if data, err := mindctrl.GetCurrentTab().Execute(transport); err != nil {
					return errors.WrapExecutionError(err, "cannot identify current tab")
				} else {
					tab = data.Id
					url = args[0]
					operation.SetTabId(tab)
					operation.SetUrl(url)
				}
			}

			if flags.Changed(REPLACE) {
				replace, _ := flags.GetBool(REPLACE)
				operation.SetReplace(true, replace)
			}

			if data, err := operation.Execute(transport); err != nil {
				return errors.WrapExecutionError(err, "cannot load url %s into tab %d", url, tab)
			} else {
				printOperationResult(stdout, tab, "loaded", data)
				return nil
			}
		}
	}
}
