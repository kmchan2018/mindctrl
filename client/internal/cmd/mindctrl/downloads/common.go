package downloads

import (
	"fmt"
	"github.com/kmchan2018/mindctrl/client/protocol"
	"io"
)

func printOperationResult(writer io.Writer, download int, operation string, data *protocol.Download) {
	fmt.Fprintf(writer, "Download %d %s.", download, operation)

	if data != nil {
		fmt.Fprintf(writer, " Details:\n\n")
		fmt.Fprintf(writer, ">> ID: %d\n", data.Id)
		fmt.Fprintf(writer, ">> URL: %s\n", data.Url)
		fmt.Fprintf(writer, ">> Filename: %s\n", data.Filename)
		fmt.Fprintf(writer, ">> Referrer: %s\n", data.Referrer)
		fmt.Fprintf(writer, ">> Mime: %s\n", data.Mime)
		fmt.Fprintf(writer, ">> Filesize: %d\n", data.Filesize)
		fmt.Fprintf(writer, ">> State: %s\n", data.State)
		fmt.Fprintf(writer, ">> Paused: %t\n", data.Paused)
		fmt.Fprintf(writer, ">> Can Resume: %t\n", data.CanResume)
		fmt.Fprintf(writer, ">> Start Time: %s\n", data.StartTime)
		fmt.Fprintf(writer, ">> Total Bytes: %d\n", data.TotalBytes)
		fmt.Fprintf(writer, ">> Received Bytes: %d\n", data.ReceivedBytes)
		fmt.Fprintf(writer, "\n")
	} else {
		fmt.Fprintf(writer, "\n\n")
	}
}
