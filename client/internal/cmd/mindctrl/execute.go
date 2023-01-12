package mindctrl

import (
	"errors"
	"fmt"
	myerrors "github.com/kmchan2018/mindctrl/client/internal/cmd/mindctrl/errors"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func Execute() {
	cobra.EnablePrefixMatching = true

	if cmd, err := RootCommand.ExecuteC(); err != nil {
		stderr := RootCommand.ErrOrStderr()

		if _, ok := err.(*myerrors.ExecutionError); ok {
			fmt.Fprintf(stderr, "ERROR: Program cannot continue due to execution error. Traceback:\n\n")

			for {
				fmt.Fprintf(stderr, "- %s\n", err.Error())
				if err = errors.Unwrap(err); err == nil {
					fmt.Fprintf(stderr, "\n")
					os.Exit(3)
				}
			}
		} else {
			fmt.Fprintf(stderr, "ERROR: %s\n\n", strings.TrimSpace(err.Error()))
			cmd.Usage()
			os.Exit(2)
		}
	}
}
