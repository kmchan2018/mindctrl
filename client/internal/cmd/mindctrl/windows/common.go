package windows

import (
	"fmt"
	"github.com/kmchan2018/mindctrl/client/protocol"
	"io"
)

func printOperationResult(writer io.Writer, window int, operation string, data *protocol.Window) {
	fmt.Fprintf(writer, "Window %d %s.", window, operation)

	if data != nil {
		fmt.Fprintf(writer, " Details:\n\n")
		fmt.Fprintf(writer, ">> ID: %d\n", data.Id)
		fmt.Fprintf(writer, ">> Type: %s\n", data.Type)
		fmt.Fprintf(writer, ">> Position: %d, %d\n", data.Left, data.Top)
		fmt.Fprintf(writer, ">> Dimension: %d x %d\n", data.Width, data.Height)
		fmt.Fprintf(writer, ">> State: %s\n", data.State)
		fmt.Fprintf(writer, ">> Focused: %t\n", data.Focused)
		fmt.Fprintf(writer, ">> Always On Top: %t\n", data.AlwaysOnTop)
		fmt.Fprintf(writer, "\n")
	} else {
		fmt.Fprintf(writer, "\n\n")
	}
}
