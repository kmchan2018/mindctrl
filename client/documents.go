package mindctrl

import (
	"errors"
	"github.com/kmchan2018/mindctrl/client/protocol"
)

// This operation provides a fluent interface to execute documents.query
// method on a mindctrl web extension instance.
//
type QueryDocumentOperation struct {
	GenericOperation
	input  protocol.QueryDocumentInput
	output protocol.QueryDocumentOutput
}

func QueryDocument(tabId int, query string, result interface{}) *QueryDocumentOperation {
	op := &QueryDocumentOperation{}
	op.input.TabId = tabId
	op.input.Query = query
	op.input.Operation = ""
	op.input.Variables = nil
	op.output.Result = result
	return op
}

func (op *QueryDocumentOperation) TabId() int {
	return op.input.TabId
}

func (op *QueryDocumentOperation) Query() string {
	return op.input.Query
}

func (op *QueryDocumentOperation) Operation() string {
	return op.input.Operation
}

func (op *QueryDocumentOperation) Variables() map[string]interface{} {
	if op.input.Variables == nil {
		return nil
	} else {
		output := make(map[string]interface{})

		for k, v := range op.input.Variables {
			output[k] = v
		}

		return output
	}
}

func (op *QueryDocumentOperation) SetTabId(tabId int) *QueryDocumentOperation {
	op.doEnsureNotStarted()
	op.input.TabId = tabId
	return op
}

func (op *QueryDocumentOperation) SetQuery(query string) *QueryDocumentOperation {
	op.doEnsureNotStarted()
	op.input.Query = query
	return op
}

func (op *QueryDocumentOperation) SetResult(result interface{}) *QueryDocumentOperation {
	op.doEnsureNotStarted()
	op.output.Result = result
	return op
}

func (op *QueryDocumentOperation) SetOperation(operation string) *QueryDocumentOperation {
	op.doEnsureNotStarted()
	op.input.Operation = operation
	return op
}

func (op *QueryDocumentOperation) SetVariable(name string, value interface{}) *QueryDocumentOperation {
	op.doEnsureNotStarted()

	if op.input.Variables != nil {
		op.input.Variables[name] = value
		return op
	} else {
		op.input.Variables = make(map[string]interface{})
		op.input.Variables[name] = value
		return op
	}
}

func (op *QueryDocumentOperation) UndefineVariable(name string) *QueryDocumentOperation {
	op.doEnsureNotStarted()

	if op.input.Variables == nil {
		return op
	} else if _, found := op.input.Variables[name]; found == false {
		return op
	} else if len(op.input.Variables) > 1 {
		delete(op.input.Variables, name)
		return op
	} else {
		op.input.Variables = nil
		return op
	}
}

func (op *QueryDocumentOperation) MergeVariables(variables map[string]interface{}) *QueryDocumentOperation {
	op.doEnsureNotStarted()

	for k, v := range variables {
		if op.input.Variables == nil {
			op.input.Variables = make(map[string]interface{})
			op.input.Variables[k] = v
		} else {
			op.input.Variables[k] = v
		}
	}

	return op
}

func (op *QueryDocumentOperation) ClearVariables() *QueryDocumentOperation {
	op.doEnsureNotStarted()
	op.input.Variables = nil
	return op
}

func (op *QueryDocumentOperation) Start(transport *Transport, callback func(op *QueryDocumentOperation)) {
	op.doStart(transport, protocol.QueryDocumentMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *QueryDocumentOperation) StartChannel(transport *Transport, channel chan *QueryDocumentOperation) {
	op.doStart(transport, protocol.QueryDocumentMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *QueryDocumentOperation) Execute(transport *Transport) error {
	if err := op.doExecute(transport, protocol.QueryDocumentMethod, &op.input, &op.output); err != nil {
		return err
	} else if op.output.Success == false {
		return errors.New(op.output.Message)
	} else {
		return nil
	}
}

func (op *QueryDocumentOperation) Result() error {
	op.doEnsureFinished()

	if op.output.Success == false {
		return errors.New(op.output.Message)
	} else {
		return nil
	}
}
