package tabs

import (
	"github.com/kmchan2018/mindctrl/client"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/errors"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/options"
	"github.com/spf13/cobra"
)

var (
	ReloadCommand = &cobra.Command{
		Use:   "reload [ tab ]",
		Short: "Reload the target tab",
		Long:  "Reload the target tab",

		DisableAutoGenTag:     true,
		DisableFlagsInUseLine: true,
	}
)

func init() {
	const BYPASS_CACHE = "bypass-cache"

	flags := ReloadCommand.Flags()
	flags.Bool(BYPASS_CACHE, false, "bypass browser cache for the reload")

	ReloadCommand.Args = func(cmd *cobra.Command, args []string) error {
		length := len(args)

		if length > 1 {
			return errors.NewExcessArgumentError()
		} else if length == 1 && options.IsId(args[0]) == false {
			return errors.NewInvalidArgumentError("tab", "argument should be a valid tab id")
		} else {
			return nil
		}
	}

	ReloadCommand.RunE = func(cmd *cobra.Command, args []string) error {
		if transport, err := options.GetTransport(cmd); err != nil {
			return errors.WrapExecutionError(err, "cannot connect to browser")
		} else {
			tab := 0
			operation := mindctrl.ReloadTab(tab)
			stdout := cmd.OutOrStdout()
			flags := cmd.Flags()

			if len(args) > 0 {
				tab = options.ParseId(args[0])
				operation.SetTabId(tab)
			} else {
				if data, err := mindctrl.GetCurrentTab().Execute(transport); err != nil {
					return errors.WrapExecutionError(err, "cannot identify current tab")
				} else {
					tab = data.Id
					operation.SetTabId(tab)
				}
			}

			if flags.Changed(BYPASS_CACHE) {
				bypassCache, _ := flags.GetBool(BYPASS_CACHE)
				operation.SetBypassCache(true, bypassCache)
			}

			if data, err := operation.Execute(transport); err != nil {
				return errors.WrapExecutionError(err, "cannot reload tab %d", tab)
			} else {
				printOperationResult(stdout, tab, "reloaded", data)
				return nil
			}
		}
	}
}
