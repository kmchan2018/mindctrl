package options

import (
	"fmt"
	"github.com/kmchan2018/mindctrl/client"
	"github.com/spf13/cobra"
	"os"
	"time"
)

func GetTransport(cmd *cobra.Command) (*mindctrl.Transport, error) {
	flags := cmd.Flags()
	server, _ := flags.GetString("server")
	browser, _ := flags.GetString("browser")
	username, _ := flags.GetString("username")
	password, _ := flags.GetString("password")

	pid := os.Getpid()
	now := time.Now().UnixMilli()
	name := fmt.Sprintf("mindctrl_golang_%d_%d", pid, now)

	if username != "" && password != "" {
		options := &mindctrl.Options{}
		options.Username = username
		options.Password = password
		return mindctrl.NewTransport(server, name, browser, options)
	} else {
		return mindctrl.NewTransport(server, name, browser, nil)
	}
}
