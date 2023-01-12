package tabs

import (
	"fmt"
	"github.com/kmchan2018/mindctrl/client/protocol"
	"io"
)

func printOperationResult(writer io.Writer, tab int, operation string, data *protocol.Tab) {
	fmt.Fprintf(writer, "Tab %d %s.", tab, operation)

	if data != nil {
		fmt.Fprintf(writer, " Details:\n\n")
		fmt.Fprintf(writer, ">> ID: %d\n", data.Id)
		fmt.Fprintf(writer, ">> Window: %d\n", data.WindowId)
		fmt.Fprintf(writer, ">> Position: %d\n", data.Index)
		fmt.Fprintf(writer, ">> Status: %s\n", data.Status)
		fmt.Fprintf(writer, ">> Active: %t\n", data.Active)
		fmt.Fprintf(writer, ">> Highlighted: %t\n", data.Highlighted)
		fmt.Fprintf(writer, ">> Pinned: %t\n", data.Pinned)
		fmt.Fprintf(writer, ">> Hidden: %t\n", data.Hidden)
		fmt.Fprintf(writer, ">> Discardable: %t\n", data.Discardable)
		fmt.Fprintf(writer, ">> Discarded: %t\n", data.Discarded)
		fmt.Fprintf(writer, ">> Attention: %t\n", data.Attention)
		fmt.Fprintf(writer, ">> Audible: %t\n", data.Audible)
		fmt.Fprintf(writer, ">> Muted: %t\n", data.Muted.Muted)
		fmt.Fprintf(writer, ">> Url: %s\n", data.Url)
		fmt.Fprintf(writer, ">> Title: %s\n", data.Title)
		fmt.Fprintf(writer, ">> Icon: %s\n", data.FavIcon)
		fmt.Fprintf(writer, "\n")
	} else {
		fmt.Fprintf(writer, "\n\n")
	}
}
