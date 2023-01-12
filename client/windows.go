package mindctrl

import (
	"errors"
	"github.com/kmchan2018/mindctrl/client/protocol"
)

// This operation provides a fluent interface to execute windows.find
// method on a mindctrl web extension instance.
//
type FindWindowsOperation struct {
	GenericOperation
	input  protocol.FindWindowsInput
	output protocol.FindWindowsOutput
}

func FindWindows() *FindWindowsOperation {
	op := &FindWindowsOperation{}
	return op
}

func (op *FindWindowsOperation) Start(transport *Transport, callback func(op *FindWindowsOperation)) {
	op.doStart(transport, protocol.FindWindowsMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *FindWindowsOperation) StartChannel(transport *Transport, channel chan *FindWindowsOperation) {
	op.doStart(transport, protocol.FindWindowsMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *FindWindowsOperation) Execute(transport *Transport) ([]protocol.Window, error) {
	if err := op.doExecute(transport, protocol.FindWindowsMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return op.output.Result, nil
	}
}

func (op *FindWindowsOperation) Result() ([]protocol.Window, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute windows.get
// method on a mindctrl web extension instance.
//
type GetWindowOperation struct {
	GenericOperation
	input  protocol.GetWindowInput
	output protocol.GetWindowOutput
}

func GetWindow(windowId int) *GetWindowOperation {
	op := &GetWindowOperation{}
	op.input.WindowId = windowId
	return op
}

func (op *GetWindowOperation) WindowId() int {
	return op.input.WindowId
}

func (op *GetWindowOperation) SetWindowId(windowId int) *GetWindowOperation {
	op.doEnsureNotStarted()
	op.input.WindowId = windowId
	return op
}

func (op *GetWindowOperation) Start(transport *Transport, callback func(op *GetWindowOperation)) {
	op.doStart(transport, protocol.GetWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *GetWindowOperation) StartChannel(transport *Transport, channel chan *GetWindowOperation) {
	op.doStart(transport, protocol.GetWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *GetWindowOperation) Execute(transport *Transport) (*protocol.Window, error) {
	if err := op.doExecute(transport, protocol.GetWindowMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *GetWindowOperation) Result() (*protocol.Window, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute windows.get_current
// method on a mindctrl web extension instance.
//
type GetCurrentWindowOperation struct {
	GenericOperation
	input  protocol.GetCurrentWindowInput
	output protocol.GetCurrentWindowOutput
}

func GetCurrentWindow() *GetCurrentWindowOperation {
	op := &GetCurrentWindowOperation{}
	return op
}

func (op *GetCurrentWindowOperation) Start(transport *Transport, callback func(op *GetCurrentWindowOperation)) {
	op.doStart(transport, protocol.GetCurrentWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *GetCurrentWindowOperation) StartChannel(transport *Transport, channel chan *GetCurrentWindowOperation) {
	op.doStart(transport, protocol.GetCurrentWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *GetCurrentWindowOperation) Execute(transport *Transport) (*protocol.Window, error) {
	if err := op.doExecute(transport, protocol.GetCurrentWindowMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *GetCurrentWindowOperation) Result() (*protocol.Window, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute windows.create
// method on a mindctrl web extension instance.
//
type CreateWindowOperation struct {
	GenericOperation
	input  protocol.CreateWindowInput
	output protocol.CreateWindowOutput
}

func CreateWindow() *CreateWindowOperation {
	op := &CreateWindowOperation{}
	return op
}

func (op *CreateWindowOperation) Start(transport *Transport, callback func(op *CreateWindowOperation)) {
	op.doStart(transport, protocol.CreateWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *CreateWindowOperation) StartChannel(transport *Transport, channel chan *CreateWindowOperation) {
	op.doStart(transport, protocol.CreateWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *CreateWindowOperation) Execute(transport *Transport) (*protocol.Window, error) {
	if err := op.doExecute(transport, protocol.CreateWindowMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *CreateWindowOperation) Result() (*protocol.Window, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute windows.move
// method on a mindctrl web extension instance.
//
type MoveWindowOperation struct {
	GenericOperation
	input  protocol.MoveWindowInput
	output protocol.MoveWindowOutput
}

func MoveWindow(windowId int, left int, top int) *MoveWindowOperation {
	op := &MoveWindowOperation{}
	op.input.WindowId = windowId
	op.input.Left = left
	op.input.Top = top
	return op
}

func (op *MoveWindowOperation) WindowId() int {
	return op.input.WindowId
}

func (op *MoveWindowOperation) Left() int {
	return op.input.Left
}

func (op *MoveWindowOperation) Top() int {
	return op.input.Top
}

func (op *MoveWindowOperation) SetWindowId(windowId int) *MoveWindowOperation {
	op.doEnsureNotStarted()
	op.input.WindowId = windowId
	return op
}

func (op *MoveWindowOperation) SetLeft(left int) *MoveWindowOperation {
	op.doEnsureNotStarted()
	op.input.Left = left
	return op
}

func (op *MoveWindowOperation) SetTop(top int) *MoveWindowOperation {
	op.doEnsureNotStarted()
	op.input.Top = top
	return op
}

func (op *MoveWindowOperation) Start(transport *Transport, callback func(op *MoveWindowOperation)) {
	op.doStart(transport, protocol.MoveWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *MoveWindowOperation) StartChannel(transport *Transport, channel chan *MoveWindowOperation) {
	op.doStart(transport, protocol.MoveWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *MoveWindowOperation) Execute(transport *Transport) (*protocol.Window, error) {
	if err := op.doExecute(transport, protocol.MoveWindowMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *MoveWindowOperation) Result() (*protocol.Window, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute windows.resize
// method on a mindctrl web extension instance.
//
type ResizeWindowOperation struct {
	GenericOperation
	input  protocol.ResizeWindowInput
	output protocol.ResizeWindowOutput
}

func ResizeWindow(windowId int, width int, height int) *ResizeWindowOperation {
	op := &ResizeWindowOperation{}
	op.input.WindowId = windowId
	op.input.Width = width
	op.input.Height = height
	return op
}

func (op *ResizeWindowOperation) WindowId() int {
	return op.input.WindowId
}

func (op *ResizeWindowOperation) Width() int {
	return op.input.Width
}

func (op *ResizeWindowOperation) Height() int {
	return op.input.Height
}

func (op *ResizeWindowOperation) SetWindowId(windowId int) *ResizeWindowOperation {
	op.doEnsureNotStarted()
	op.input.WindowId = windowId
	return op
}

func (op *ResizeWindowOperation) SetWidth(width int) *ResizeWindowOperation {
	op.doEnsureNotStarted()
	op.input.Width = width
	return op
}

func (op *ResizeWindowOperation) SetHeight(height int) *ResizeWindowOperation {
	op.doEnsureNotStarted()
	op.input.Height = height
	return op
}

func (op *ResizeWindowOperation) Start(transport *Transport, callback func(op *ResizeWindowOperation)) {
	op.doStart(transport, protocol.ResizeWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *ResizeWindowOperation) StartChannel(transport *Transport, channel chan *ResizeWindowOperation) {
	op.doStart(transport, protocol.ResizeWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *ResizeWindowOperation) Execute(transport *Transport) (*protocol.Window, error) {
	if err := op.doExecute(transport, protocol.ResizeWindowMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *ResizeWindowOperation) Result() (*protocol.Window, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute windows.minimize
// method on a mindctrl web extension instance.
//
type MinimizeWindowOperation struct {
	GenericOperation
	input  protocol.MinimizeWindowInput
	output protocol.MinimizeWindowOutput
}

func MinimizeWindow(windowId int) *MinimizeWindowOperation {
	op := &MinimizeWindowOperation{}
	op.input.WindowId = windowId
	return op
}

func (op *MinimizeWindowOperation) WindowId() int {
	return op.input.WindowId
}

func (op *MinimizeWindowOperation) SetWindowId(windowId int) *MinimizeWindowOperation {
	op.doEnsureNotStarted()
	op.input.WindowId = windowId
	return op
}

func (op *MinimizeWindowOperation) Start(transport *Transport, callback func(op *MinimizeWindowOperation)) {
	op.doStart(transport, protocol.MinimizeWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *MinimizeWindowOperation) StartChannel(transport *Transport, channel chan *MinimizeWindowOperation) {
	op.doStart(transport, protocol.MinimizeWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *MinimizeWindowOperation) Execute(transport *Transport) (*protocol.Window, error) {
	if err := op.doExecute(transport, protocol.MinimizeWindowMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *MinimizeWindowOperation) Result() (*protocol.Window, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute windows.maximize
// method on a mindctrl web extension instance.
//
type MaximizeWindowOperation struct {
	GenericOperation
	input  protocol.MaximizeWindowInput
	output protocol.MaximizeWindowOutput
}

func MaximizeWindow(windowId int) *MaximizeWindowOperation {
	op := &MaximizeWindowOperation{}
	op.input.WindowId = windowId
	return op
}

func (op *MaximizeWindowOperation) WindowId() int {
	return op.input.WindowId
}

func (op *MaximizeWindowOperation) SetWindowId(windowId int) *MaximizeWindowOperation {
	op.doEnsureNotStarted()
	op.input.WindowId = windowId
	return op
}

func (op *MaximizeWindowOperation) Start(transport *Transport, callback func(op *MaximizeWindowOperation)) {
	op.doStart(transport, protocol.MaximizeWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *MaximizeWindowOperation) StartChannel(transport *Transport, channel chan *MaximizeWindowOperation) {
	op.doStart(transport, protocol.MaximizeWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *MaximizeWindowOperation) Execute(transport *Transport) (*protocol.Window, error) {
	if err := op.doExecute(transport, protocol.MaximizeWindowMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *MaximizeWindowOperation) Result() (*protocol.Window, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute windows.fullscreen
// method on a mindctrl web extension instance.
//
type FullscreenWindowOperation struct {
	GenericOperation
	input  protocol.FullscreenWindowInput
	output protocol.FullscreenWindowOutput
}

func FullscreenWindow(windowId int) *FullscreenWindowOperation {
	op := &FullscreenWindowOperation{}
	op.input.WindowId = windowId
	return op
}

func (op *FullscreenWindowOperation) WindowId() int {
	return op.input.WindowId
}

func (op *FullscreenWindowOperation) SetWindowId(windowId int) *FullscreenWindowOperation {
	op.doEnsureNotStarted()
	op.input.WindowId = windowId
	return op
}

func (op *FullscreenWindowOperation) Start(transport *Transport, callback func(op *FullscreenWindowOperation)) {
	op.doStart(transport, protocol.FullscreenWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *FullscreenWindowOperation) StartChannel(transport *Transport, channel chan *FullscreenWindowOperation) {
	op.doStart(transport, protocol.FullscreenWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *FullscreenWindowOperation) Execute(transport *Transport) (*protocol.Window, error) {
	if err := op.doExecute(transport, protocol.FullscreenWindowMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *FullscreenWindowOperation) Result() (*protocol.Window, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute windows.restore
// method on a mindctrl web extension instance.
//
type RestoreWindowOperation struct {
	GenericOperation
	input  protocol.RestoreWindowInput
	output protocol.RestoreWindowOutput
}

func RestoreWindow(windowId int) *RestoreWindowOperation {
	op := &RestoreWindowOperation{}
	op.input.WindowId = windowId
	return op
}

func (op *RestoreWindowOperation) WindowId() int {
	return op.input.WindowId
}

func (op *RestoreWindowOperation) SetWindowId(windowId int) *RestoreWindowOperation {
	op.doEnsureNotStarted()
	op.input.WindowId = windowId
	return op
}

func (op *RestoreWindowOperation) Start(transport *Transport, callback func(op *RestoreWindowOperation)) {
	op.doStart(transport, protocol.RestoreWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *RestoreWindowOperation) StartChannel(transport *Transport, channel chan *RestoreWindowOperation) {
	op.doStart(transport, protocol.RestoreWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *RestoreWindowOperation) Execute(transport *Transport) (*protocol.Window, error) {
	if err := op.doExecute(transport, protocol.RestoreWindowMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *RestoreWindowOperation) Result() (*protocol.Window, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute windows.focus
// method on a mindctrl web extension instance.
//
type FocusWindowOperation struct {
	GenericOperation
	input  protocol.FocusWindowInput
	output protocol.FocusWindowOutput
}

func FocusWindow(windowId int) *FocusWindowOperation {
	op := &FocusWindowOperation{}
	op.input.WindowId = windowId
	return op
}

func (op *FocusWindowOperation) WindowId() int {
	return op.input.WindowId
}

func (op *FocusWindowOperation) SetWindowId(windowId int) *FocusWindowOperation {
	op.doEnsureNotStarted()
	op.input.WindowId = windowId
	return op
}

func (op *FocusWindowOperation) Start(transport *Transport, callback func(op *FocusWindowOperation)) {
	op.doStart(transport, protocol.FocusWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *FocusWindowOperation) StartChannel(transport *Transport, channel chan *FocusWindowOperation) {
	op.doStart(transport, protocol.FocusWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *FocusWindowOperation) Execute(transport *Transport) (*protocol.Window, error) {
	if err := op.doExecute(transport, protocol.FocusWindowMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *FocusWindowOperation) Result() (*protocol.Window, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute windows.unfocus
// method on a mindctrl web extension instance.
//
type UnfocusWindowOperation struct {
	GenericOperation
	input  protocol.UnfocusWindowInput
	output protocol.UnfocusWindowOutput
}

func UnfocusWindow(windowId int) *UnfocusWindowOperation {
	op := &UnfocusWindowOperation{}
	op.input.WindowId = windowId
	return op
}

func (op *UnfocusWindowOperation) WindowId() int {
	return op.input.WindowId
}

func (op *UnfocusWindowOperation) SetWindowId(windowId int) *UnfocusWindowOperation {
	op.doEnsureNotStarted()
	op.input.WindowId = windowId
	return op
}

func (op *UnfocusWindowOperation) Start(transport *Transport, callback func(op *UnfocusWindowOperation)) {
	op.doStart(transport, protocol.UnfocusWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *UnfocusWindowOperation) StartChannel(transport *Transport, channel chan *UnfocusWindowOperation) {
	op.doStart(transport, protocol.UnfocusWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *UnfocusWindowOperation) Execute(transport *Transport) (*protocol.Window, error) {
	if err := op.doExecute(transport, protocol.UnfocusWindowMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *UnfocusWindowOperation) Result() (*protocol.Window, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute windows.remove
// method on a mindctrl web extension instance.
//
type RemoveWindowOperation struct {
	GenericOperation
	input  protocol.RemoveWindowInput
	output protocol.RemoveWindowOutput
}

func RemoveWindow(windowId int) *RemoveWindowOperation {
	op := &RemoveWindowOperation{}
	op.input.WindowId = windowId
	return op
}

func (op *RemoveWindowOperation) WindowId() int {
	return op.input.WindowId
}

func (op *RemoveWindowOperation) SetWindowId(windowId int) *RemoveWindowOperation {
	op.doEnsureNotStarted()
	op.input.WindowId = windowId
	return op
}

func (op *RemoveWindowOperation) Start(transport *Transport, callback func(op *RemoveWindowOperation)) {
	op.doStart(transport, protocol.RemoveWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *RemoveWindowOperation) StartChannel(transport *Transport, channel chan *RemoveWindowOperation) {
	op.doStart(transport, protocol.RemoveWindowMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *RemoveWindowOperation) Execute(transport *Transport) error {
	if err := op.doExecute(transport, protocol.RemoveWindowMethod, &op.input, &op.output); err != nil {
		return err
	} else if op.output.Success == false {
		return errors.New(op.output.Message)
	} else {
		return nil
	}
}

func (op *RemoveWindowOperation) Result() error {
	op.doEnsureFinished()

	if op.output.Success == false {
		return errors.New(op.output.Message)
	} else {
		return nil
	}
}
