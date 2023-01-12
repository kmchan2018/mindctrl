package mindctrl

import (
	"errors"
	"github.com/kmchan2018/mindctrl/client/protocol"
)

// This operation provides a fluent interface to execute downloads.find
// method on a mindctrl web extension instance.
//
type FindDownloadsOperation struct {
	GenericOperation
	input  protocol.FindDownloadsInput
	output protocol.FindDownloadsOutput
}

func FindDownloads() *FindDownloadsOperation {
	op := &FindDownloadsOperation{}
	op.input.Url = nil
	op.input.State = nil
	return op
}

func (op *FindDownloadsOperation) Url() (bool, string) {
	if op.input.Url != nil {
		return true, op.input.UrlValue
	} else {
		return false, ""
	}
}

func (op *FindDownloadsOperation) State() (bool, string) {
	if op.input.State != nil {
		return true, op.input.StateValue
	} else {
		return false, ""
	}
}

func (op *FindDownloadsOperation) SetUrl(specified bool, url string) *FindDownloadsOperation {
	op.doEnsureNotStarted()

	if specified {
		op.input.UrlValue = url
		op.input.Url = &op.input.UrlValue
		return op
	} else {
		op.input.Url = nil
		return op
	}
}

func (op *FindDownloadsOperation) SetState(specified bool, state string) *FindDownloadsOperation {
	op.doEnsureNotStarted()

	if specified {
		op.input.StateValue = state
		op.input.State = &op.input.StateValue
		return op
	} else {
		op.input.State = nil
		return op
	}
}

func (op *FindDownloadsOperation) Start(transport *Transport, callback func(op *FindDownloadsOperation)) {
	op.doStart(transport, protocol.FindDownloadsMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *FindDownloadsOperation) StartChannel(transport *Transport, channel chan *FindDownloadsOperation) {
	op.doStart(transport, protocol.FindDownloadsMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *FindDownloadsOperation) Execute(transport *Transport) ([]protocol.Download, error) {
	if err := op.doExecute(transport, protocol.FindDownloadsMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return op.output.Result, nil
	}
}

func (op *FindDownloadsOperation) Result() ([]protocol.Download, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute downloads.get
// method on a mindctrl web extension instance.
//
type GetDownloadOperation struct {
	GenericOperation
	input  protocol.GetDownloadInput
	output protocol.GetDownloadOutput
}

func GetDownload(downloadId int) *GetDownloadOperation {
	op := &GetDownloadOperation{}
	op.input.DownloadId = downloadId
	return op
}

func (op *GetDownloadOperation) DownloadId() int {
	return op.input.DownloadId
}

func (op *GetDownloadOperation) SetDownloadId(downloadId int) *GetDownloadOperation {
	op.doEnsureNotStarted()
	op.input.DownloadId = downloadId
	return op
}

func (op *GetDownloadOperation) Start(transport *Transport, callback func(op *GetDownloadOperation)) {
	op.doStart(transport, protocol.GetDownloadMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *GetDownloadOperation) StartChannel(transport *Transport, channel chan *GetDownloadOperation) {
	op.doStart(transport, protocol.GetDownloadMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *GetDownloadOperation) Execute(transport *Transport) (*protocol.Download, error) {
	if err := op.doExecute(transport, protocol.GetDownloadMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *GetDownloadOperation) Result() (*protocol.Download, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute downloads.create
// method on a mindctrl web extension instance.
//
type CreateDownloadOperation struct {
	GenericOperation
	input  protocol.CreateDownloadInput
	output protocol.CreateDownloadOutput
}

func CreateDownload(url string, filename string) *CreateDownloadOperation {
	op := &CreateDownloadOperation{}
	op.input.Url = url
	op.input.Filename = filename
	op.input.Referrer = nil
	op.input.NoWait = nil
	return op
}

func (op *CreateDownloadOperation) Url() string {
	return op.input.Url
}

func (op *CreateDownloadOperation) Filename() string {
	return op.input.Filename
}

func (op *CreateDownloadOperation) Referrer() (bool, string) {
	if op.input.Referrer != nil {
		return true, op.input.ReferrerValue
	} else {
		return false, ""
	}
}

func (op *CreateDownloadOperation) NoWait() (bool, bool) {
	if op.input.NoWait != nil {
		return true, op.input.NoWaitValue
	} else {
		return false, false
	}
}

func (op *CreateDownloadOperation) SetUrl(url string) *CreateDownloadOperation {
	op.doEnsureNotStarted()
	op.input.Url = url
	return op
}

func (op *CreateDownloadOperation) SetFilename(filename string) *CreateDownloadOperation {
	op.doEnsureNotStarted()
	op.input.Filename = filename
	return op
}

func (op *CreateDownloadOperation) SetReferrer(specified bool, referrer string) *CreateDownloadOperation {
	op.doEnsureNotStarted()

	if specified {
		op.input.ReferrerValue = referrer
		op.input.Referrer = &op.input.ReferrerValue
		return op
	} else {
		op.input.Referrer = nil
		return op
	}
}

func (op *CreateDownloadOperation) SetNoWait(specified bool, noWait bool) *CreateDownloadOperation {
	op.doEnsureNotStarted()

	if specified {
		op.input.NoWaitValue = noWait
		op.input.NoWait = &op.input.NoWaitValue
		return op
	} else {
		op.input.NoWait = nil
		return op
	}
}

func (op *CreateDownloadOperation) Start(transport *Transport, callback func(op *CreateDownloadOperation)) {
	op.doStart(transport, protocol.CreateDownloadMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *CreateDownloadOperation) StartChannel(transport *Transport, channel chan *CreateDownloadOperation) {
	op.doStart(transport, protocol.CreateDownloadMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *CreateDownloadOperation) Execute(transport *Transport) (*protocol.Download, error) {
	if err := op.doExecute(transport, protocol.CreateDownloadMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *CreateDownloadOperation) Result() (*protocol.Download, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute downloads.pause
// method on a mindctrl web extension instance.
//
type PauseDownloadOperation struct {
	GenericOperation
	input  protocol.PauseDownloadInput
	output protocol.PauseDownloadOutput
}

func PauseDownload(downloadId int) *PauseDownloadOperation {
	op := &PauseDownloadOperation{}
	op.input.DownloadId = downloadId
	return op
}

func (op *PauseDownloadOperation) DownloadId() int {
	return op.input.DownloadId
}

func (op *PauseDownloadOperation) SetDownloadId(downloadId int) *PauseDownloadOperation {
	op.doEnsureNotStarted()
	op.input.DownloadId = downloadId
	return op
}

func (op *PauseDownloadOperation) Start(transport *Transport, callback func(op *PauseDownloadOperation)) {
	op.doStart(transport, protocol.PauseDownloadMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *PauseDownloadOperation) StartChannel(transport *Transport, channel chan *PauseDownloadOperation) {
	op.doStart(transport, protocol.PauseDownloadMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *PauseDownloadOperation) Execute(transport *Transport) (*protocol.Download, error) {
	if err := op.doExecute(transport, protocol.PauseDownloadMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *PauseDownloadOperation) Result() (*protocol.Download, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute downloads.resume
// method on a mindctrl web extension instance.
//
type ResumeDownloadOperation struct {
	GenericOperation
	input  protocol.ResumeDownloadInput
	output protocol.ResumeDownloadOutput
}

func ResumeDownload(downloadId int) *ResumeDownloadOperation {
	op := &ResumeDownloadOperation{}
	op.input.DownloadId = downloadId
	op.input.NoWait = nil
	return op
}

func (op *ResumeDownloadOperation) DownloadId() int {
	return op.input.DownloadId
}

func (op *ResumeDownloadOperation) NoWait() (bool, bool) {
	if op.input.NoWait != nil {
		return true, op.input.NoWaitValue
	} else {
		return false, false
	}
}

func (op *ResumeDownloadOperation) SetDownloadId(downloadId int) *ResumeDownloadOperation {
	op.doEnsureNotStarted()
	op.input.DownloadId = downloadId
	return op
}

func (op *ResumeDownloadOperation) SetNoWait(specified bool, noWait bool) *ResumeDownloadOperation {
	op.doEnsureNotStarted()

	if specified {
		op.input.NoWaitValue = noWait
		op.input.NoWait = &op.input.NoWaitValue
		return op
	} else {
		op.input.NoWait = nil
		return op
	}
}

func (op *ResumeDownloadOperation) Start(transport *Transport, callback func(op *ResumeDownloadOperation)) {
	op.doStart(transport, protocol.ResumeDownloadMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *ResumeDownloadOperation) StartChannel(transport *Transport, channel chan *ResumeDownloadOperation) {
	op.doStart(transport, protocol.ResumeDownloadMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *ResumeDownloadOperation) Execute(transport *Transport) (*protocol.Download, error) {
	if err := op.doExecute(transport, protocol.ResumeDownloadMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *ResumeDownloadOperation) Result() (*protocol.Download, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute downloads.cancel
// method on a mindctrl web extension instance.
//
type CancelDownloadOperation struct {
	GenericOperation
	input  protocol.CancelDownloadInput
	output protocol.CancelDownloadOutput
}

func CancelDownload(downloadId int) *CancelDownloadOperation {
	op := &CancelDownloadOperation{}
	op.input.DownloadId = downloadId
	return op
}

func (op *CancelDownloadOperation) DownloadId() int {
	return op.input.DownloadId
}

func (op *CancelDownloadOperation) SetDownloadId(downloadId int) *CancelDownloadOperation {
	op.doEnsureNotStarted()
	op.input.DownloadId = downloadId
	return op
}

func (op *CancelDownloadOperation) Start(transport *Transport, callback func(op *CancelDownloadOperation)) {
	op.doStart(transport, protocol.CancelDownloadMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *CancelDownloadOperation) StartChannel(transport *Transport, channel chan *CancelDownloadOperation) {
	op.doStart(transport, protocol.CancelDownloadMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *CancelDownloadOperation) Execute(transport *Transport) (*protocol.Download, error) {
	if err := op.doExecute(transport, protocol.CancelDownloadMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *CancelDownloadOperation) Result() (*protocol.Download, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute downloads.remove
// method on a mindctrl web extension instance.
//
type RemoveDownloadOperation struct {
	GenericOperation
	input  protocol.RemoveDownloadInput
	output protocol.RemoveDownloadOutput
}

func RemoveDownload(downloadId int) *RemoveDownloadOperation {
	op := &RemoveDownloadOperation{}
	op.input.DownloadId = downloadId
	return op
}

func (op *RemoveDownloadOperation) DownloadId() int {
	return op.input.DownloadId
}

func (op *RemoveDownloadOperation) SetDownloadId(downloadId int) *RemoveDownloadOperation {
	op.doEnsureNotStarted()
	op.input.DownloadId = downloadId
	return op
}

func (op *RemoveDownloadOperation) Start(transport *Transport, callback func(op *RemoveDownloadOperation)) {
	op.doStart(transport, protocol.RemoveDownloadMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *RemoveDownloadOperation) StartChannel(transport *Transport, channel chan *RemoveDownloadOperation) {
	op.doStart(transport, protocol.RemoveDownloadMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *RemoveDownloadOperation) Execute(transport *Transport) error {
	if err := op.doExecute(transport, protocol.RemoveDownloadMethod, &op.input, &op.output); err != nil {
		return err
	} else if op.output.Success == false {
		return errors.New(op.output.Message)
	} else {
		return nil
	}
}

func (op *RemoveDownloadOperation) Result() error {
	op.doEnsureFinished()

	if op.output.Success == false {
		return errors.New(op.output.Message)
	} else {
		return nil
	}
}
