package mindctrl

import (
	"errors"
	"github.com/kmchan2018/mindctrl/client/protocol"
)

// This operation provides a fluent interface to execute ping method
// on a mindctrl web extension instance.
//
type PingOperation struct {
	GenericOperation
	input  protocol.PingInput
	output protocol.PingOutput
}

func Ping() *PingOperation {
	op := &PingOperation{}
	return op
}

func (op *PingOperation) Start(transport *Transport, callback func(op *PingOperation)) {
	op.doStart(transport, protocol.PingMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *PingOperation) StartChannel(transport *Transport, channel chan *PingOperation) {
	op.doStart(transport, protocol.PingMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *PingOperation) Execute(transport *Transport) error {
	if err := op.doExecute(transport, protocol.PingMethod, &op.input, &op.output); err != nil {
		return err
	} else if op.output.Success == false {
		return errors.New(op.output.Message)
	} else {
		return nil
	}
}

func (op *PingOperation) Result() error {
	op.doEnsureFinished()

	if op.output.Success == false {
		return errors.New(op.output.Message)
	} else {
		return nil
	}
}
