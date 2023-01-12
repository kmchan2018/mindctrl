package mindctrl

import (
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/downloads"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/errors"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/info"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/tabs"
	"github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/windows"
	"github.com/spf13/cobra"
	"os"
)

var (
	RootCommand = &cobra.Command{
		Use:   "mindctrl",
		Short: "Mind control your browser for fun and profit",
		Long:  "Mind control your browser for fun and profit",

		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func init() {
	RootCommand.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		return errors.NewFlagError(err)
	})

	RootCommand.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		server, _ := cmd.Flags().GetString("server")
		browser, _ := cmd.Flags().GetString("browser")

		if server == "" {
			return errors.NewArgumentError("unknown url to intermediate mqtt server")
		} else if browser == "" {
			return errors.NewArgumentError("unknown browser name")
		} else {
			return nil
		}
	}

	SERVER := os.Getenv("MINDCTRL_SERVER")
	BROWSER := os.Getenv("MINDCTRL_BROWSER")
	USERNAME := os.Getenv("MINDCTRL_USERNAME")
	PASSWORD := os.Getenv("MINDCTRL_PASSWORD")

	RootCommand.PersistentFlags().StringP("server", "s", SERVER, "url to the intermediate MQTT server")
	RootCommand.PersistentFlags().StringP("browser", "b", BROWSER, "name for the browser")
	RootCommand.PersistentFlags().StringP("username", "u", USERNAME, "username for the intermediate MQTT server")
	RootCommand.PersistentFlags().StringP("password", "p", PASSWORD, "password for the intermediate MQTT server")
	RootCommand.MarkFlagsRequiredTogether("username", "password")

	RootCommand.SetUsageTemplate(RootCommand.UsageTemplate() + "\n")
	RootCommand.AddCommand(downloads.RootCommand)
	RootCommand.AddCommand(info.RootCommand)
	RootCommand.AddCommand(tabs.RootCommand)
	RootCommand.AddCommand(windows.RootCommand)
}
