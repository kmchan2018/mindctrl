package mindctrl

import (
	"errors"
	"github.com/kmchan2018/mindctrl/client/protocol"
)

// This operation provides a fluent interface to execute info.get_browser
// method on a mindctrl web extension instance.
//
type GetBrowserInfoOperation struct {
	GenericOperation
	input  protocol.GetBrowserInfoInput
	output protocol.GetBrowserInfoOutput
}

func GetBrowserInfo() *GetBrowserInfoOperation {
	op := &GetBrowserInfoOperation{}
	return op
}

func (op *GetBrowserInfoOperation) Start(transport *Transport, callback func(op *GetBrowserInfoOperation)) {
	op.doStart(transport, protocol.GetBrowserInfoMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *GetBrowserInfoOperation) StartChannel(transport *Transport, channel chan *GetBrowserInfoOperation) {
	op.doStart(transport, protocol.GetBrowserInfoMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *GetBrowserInfoOperation) Execute(transport *Transport) (*protocol.BrowserInfo, error) {
	if err := op.doExecute(transport, protocol.GetBrowserInfoMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *GetBrowserInfoOperation) Result() (*protocol.BrowserInfo, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute info.get_platform
// method on a mindctrl web extension instance.
//
type GetPlatformInfoOperation struct {
	GenericOperation
	input  protocol.GetPlatformInfoInput
	output protocol.GetPlatformInfoOutput
}

func GetPlatformInfo() *GetPlatformInfoOperation {
	instance := &GetPlatformInfoOperation{}
	return instance
}

func (op *GetPlatformInfoOperation) Start(transport *Transport, callback func(op *GetPlatformInfoOperation)) {
	op.doStart(transport, protocol.GetPlatformInfoMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *GetPlatformInfoOperation) StartChannel(transport *Transport, channel chan *GetPlatformInfoOperation) {
	op.doStart(transport, protocol.GetPlatformInfoMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *GetPlatformInfoOperation) Execute(transport *Transport) (*protocol.PlatformInfo, error) {
	if err := op.doExecute(transport, protocol.GetPlatformInfoMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *GetPlatformInfoOperation) Result() (*protocol.PlatformInfo, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}
