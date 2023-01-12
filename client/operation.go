package mindctrl

// Embeddable struct that provides a partial implementation of
// operations.
//
type GenericOperation struct {
	started  bool
	finished bool
	err      error
}

// Return if the operation has started. When an operation has
// started, any changes to the operation will result in a panic.
//
func (op *GenericOperation) Started() bool {
	return op.started
}

// Return if the operation has finished. Before an operation has
// finished, retrieval of operation result will result in a panic.
//
func (op *GenericOperation) Finished() bool {
	return op.finished
}

func (op *GenericOperation) doEnsureNotStarted() {
	if op.started == true {
		panic("operation already started")
	}
}

func (op *GenericOperation) doEnsureFinished() {
	if op.started == false {
		panic("operation not started")
	} else if op.finished == false {
		panic("operation not finished")
	}
}

func (op *GenericOperation) doStart(transport *Transport, method string, arguments, reply interface{}, callback Callback) {
	if op.started == true {
		panic("operation already started")
	} else {
		op.started = true
		transport.start(method, arguments, reply, callback)
	}
}

func (op *GenericOperation) doExecute(transport *Transport, method string, arguments, reply interface{}) error {
	if op.started == true {
		panic("operation already started")
	} else {
		op.started = true
		op.err = transport.call(method, arguments, reply)
		op.finished = true
		return op.err
	}
}

func (op *GenericOperation) doFinish(err error) {
	if op.started == false {
		panic("operation not started")
	} else if op.finished == true {
		panic("operation already finished")
	} else {
		op.err = err
		op.finished = true
	}
}

func (op *GenericOperation) doGetError() error {
	return op.err
}
