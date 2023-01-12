package mindctrl

import (
	"errors"
	"github.com/kmchan2018/mindctrl/client/protocol"
)

// This operation provides a fluent interface to execute tabs.find
// method on a mindctrl web extension instance.
//
type FindTabsOperation struct {
	GenericOperation
	input  protocol.FindTabsInput
	output protocol.FindTabsOutput
}

func FindTabs() *FindTabsOperation {
	op := &FindTabsOperation{}
	op.input.WindowId = nil
	op.input.Url = nil
	op.input.Status = nil
	op.input.Active = nil
	op.input.Audible = nil
	op.input.Discarded = nil
	op.input.Muted = nil
	op.input.Pinned = nil
	return op
}

func (op *FindTabsOperation) WindowId() (bool, int) {
	if op.input.WindowId != nil {
		return true, op.input.WindowIdValue
	} else {
		return false, -2
	}
}

func (op *FindTabsOperation) Url() (bool, string) {
	if op.input.Url != nil {
		return true, op.input.UrlValue
	} else {
		return false, ""
	}
}

func (op *FindTabsOperation) Status() (bool, string) {
	if op.input.Status != nil {
		return true, op.input.StatusValue
	} else {
		return false, ""
	}
}

func (op *FindTabsOperation) Active() (bool, bool) {
	if op.input.Active != nil {
		return true, op.input.ActiveValue
	} else {
		return false, false
	}
}

func (op *FindTabsOperation) Audible() (bool, bool) {
	if op.input.Audible != nil {
		return true, op.input.AudibleValue
	} else {
		return false, false
	}
}

func (op *FindTabsOperation) Discarded() (bool, bool) {
	if op.input.Discarded != nil {
		return true, op.input.DiscardedValue
	} else {
		return false, false
	}
}

func (op *FindTabsOperation) Muted() (bool, bool) {
	if op.input.Muted != nil {
		return true, op.input.MutedValue
	} else {
		return false, false
	}
}

func (op *FindTabsOperation) Pinned() (bool, bool) {
	if op.input.Pinned != nil {
		return true, op.input.PinnedValue
	} else {
		return false, false
	}
}

func (op *FindTabsOperation) SetWindowId(specified bool, windowId int) *FindTabsOperation {
	op.doEnsureNotStarted()

	if specified {
		op.input.WindowIdValue = windowId
		op.input.WindowId = &op.input.WindowIdValue
		return op
	} else {
		op.input.WindowId = nil
		return op
	}
}

func (op *FindTabsOperation) SetUrl(specified bool, url string) *FindTabsOperation {
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

func (op *FindTabsOperation) SetStatus(specified bool, status string) *FindTabsOperation {
	op.doEnsureNotStarted()

	if specified {
		op.input.StatusValue = status
		op.input.Status = &op.input.StatusValue
		return op
	} else {
		op.input.Status = nil
		return op
	}
}

func (op *FindTabsOperation) SetActive(specified bool, active bool) *FindTabsOperation {
	op.doEnsureNotStarted()

	if specified {
		op.input.ActiveValue = active
		op.input.Active = &op.input.ActiveValue
		return op
	} else {
		op.input.Active = nil
		return op
	}
}

func (op *FindTabsOperation) SetAudible(specified bool, audible bool) *FindTabsOperation {
	op.doEnsureNotStarted()

	if specified {
		op.input.AudibleValue = audible
		op.input.Audible = &op.input.AudibleValue
		return op
	} else {
		op.input.Audible = nil
		return op
	}
}

func (op *FindTabsOperation) SetDiscarded(specified bool, discarded bool) *FindTabsOperation {
	op.doEnsureNotStarted()

	if specified {
		op.input.DiscardedValue = discarded
		op.input.Discarded = &op.input.DiscardedValue
		return op
	} else {
		op.input.Discarded = nil
		return op
	}
}

func (op *FindTabsOperation) SetMuted(specified bool, muted bool) *FindTabsOperation {
	op.doEnsureNotStarted()

	if specified {
		op.input.MutedValue = muted
		op.input.Muted = &op.input.MutedValue
		return op
	} else {
		op.input.Muted = nil
		return op
	}
}

func (op *FindTabsOperation) SetPinned(specified bool, pinned bool) *FindTabsOperation {
	op.doEnsureNotStarted()

	if specified {
		op.input.PinnedValue = pinned
		op.input.Pinned = &op.input.PinnedValue
		return op
	} else {
		op.input.Pinned = nil
		return op
	}
}

func (op *FindTabsOperation) Start(transport *Transport, callback func(op *FindTabsOperation)) {
	op.doStart(transport, protocol.FindTabsMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *FindTabsOperation) StartChannel(transport *Transport, channel chan *FindTabsOperation) {
	op.doStart(transport, protocol.FindTabsMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *FindTabsOperation) Execute(transport *Transport) ([]protocol.Tab, error) {
	if err := op.doExecute(transport, protocol.FindTabsMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return op.output.Result, nil
	}
}

func (op *FindTabsOperation) Result() ([]protocol.Tab, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute tabs.get
// method on a mindctrl web extension instance.
//
type GetTabOperation struct {
	GenericOperation
	input  protocol.GetTabInput
	output protocol.GetTabOutput
}

func GetTab(tabId int) *GetTabOperation {
	op := &GetTabOperation{}
	op.input.TabId = tabId
	return op
}

func (op *GetTabOperation) TabId() int {
	return op.input.TabId
}

func (op *GetTabOperation) SetTabId(tabId int) *GetTabOperation {
	op.doEnsureNotStarted()
	op.input.TabId = tabId
	return op
}

func (op *GetTabOperation) Start(transport *Transport, callback func(op *GetTabOperation)) {
	op.doStart(transport, protocol.GetTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *GetTabOperation) StartChannel(transport *Transport, channel chan *GetTabOperation) {
	op.doStart(transport, protocol.GetTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *GetTabOperation) Execute(transport *Transport) (*protocol.Tab, error) {
	if err := op.doExecute(transport, protocol.GetTabMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *GetTabOperation) Result() (*protocol.Tab, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute tabs.get_current
// method on a mindctrl web extension instance.
//
type GetCurrentTabOperation struct {
	GenericOperation
	input  protocol.GetCurrentTabInput
	output protocol.GetCurrentTabOutput
}

func GetCurrentTab() *GetCurrentTabOperation {
	op := &GetCurrentTabOperation{}
	return op
}

func (op *GetCurrentTabOperation) Start(transport *Transport, callback func(op *GetCurrentTabOperation)) {
	op.doStart(transport, protocol.GetCurrentTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *GetCurrentTabOperation) StartChannel(transport *Transport, channel chan *GetCurrentTabOperation) {
	op.doStart(transport, protocol.GetCurrentTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *GetCurrentTabOperation) Execute(transport *Transport) (*protocol.Tab, error) {
	if err := op.doExecute(transport, protocol.GetCurrentTabMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *GetCurrentTabOperation) Result() (*protocol.Tab, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute tabs.create
// method on a mindctrl web extension instance.
//
type CreateTabOperation struct {
	GenericOperation
	input  protocol.CreateTabInput
	output protocol.CreateTabOutput
}

func CreateTab() *CreateTabOperation {
	op := &CreateTabOperation{}
	op.input.WindowId = nil
	op.input.Active = nil
	op.input.Url = nil
	op.input.NoWait = nil
	return op
}

func (op *CreateTabOperation) WindowId() (bool, int) {
	if op.input.WindowId != nil {
		return true, op.input.WindowIdValue
	} else {
		return false, -2
	}
}

func (op *CreateTabOperation) Url() (bool, string) {
	if op.input.Url != nil {
		return true, op.input.UrlValue
	} else {
		return false, ""
	}
}

func (op *CreateTabOperation) Active() (bool, bool) {
	if op.input.Active != nil {
		return true, op.input.ActiveValue
	} else {
		return false, false
	}
}

func (op *CreateTabOperation) NoWait() (bool, bool) {
	if op.input.NoWait != nil {
		return true, op.input.NoWaitValue
	} else {
		return false, false
	}
}

func (op *CreateTabOperation) SetWindowId(specified bool, windowId int) *CreateTabOperation {
	op.doEnsureNotStarted()

	if specified {
		op.input.WindowIdValue = windowId
		op.input.WindowId = &op.input.WindowIdValue
		return op
	} else {
		op.input.WindowId = nil
		return op
	}
}

func (op *CreateTabOperation) SetUrl(specified bool, url string) *CreateTabOperation {
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

func (op *CreateTabOperation) SetActive(specified bool, active bool) *CreateTabOperation {
	op.doEnsureNotStarted()

	if specified {
		op.input.ActiveValue = active
		op.input.Active = &op.input.ActiveValue
		return op
	} else {
		op.input.Active = nil
		return op
	}
}

func (op *CreateTabOperation) SetNoWait(specified bool, noWait bool) *CreateTabOperation {
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

func (op *CreateTabOperation) Start(transport *Transport, callback func(op *CreateTabOperation)) {
	op.doStart(transport, protocol.CreateTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *CreateTabOperation) StartChannel(transport *Transport, channel chan *CreateTabOperation) {
	op.doStart(transport, protocol.CreateTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *CreateTabOperation) Execute(transport *Transport) (*protocol.Tab, error) {
	if err := op.doExecute(transport, protocol.CreateTabMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *CreateTabOperation) Result() (*protocol.Tab, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute tabs.load
// method on a mindctrl web extension instance.
//
type LoadTabOperation struct {
	GenericOperation
	input  protocol.LoadTabInput
	output protocol.LoadTabOutput
}

func LoadTab(tabId int, url string) *LoadTabOperation {
	op := &LoadTabOperation{}
	op.input.TabId = tabId
	op.input.Url = url
	op.input.Replace = nil
	op.input.NoWait = nil
	return op
}

func (op *LoadTabOperation) TabId() int {
	return op.input.TabId
}

func (op *LoadTabOperation) Url() string {
	return op.input.Url
}

func (op *LoadTabOperation) Replace() (bool, bool) {
	if op.input.Replace != nil {
		return true, op.input.ReplaceValue
	} else {
		return false, false
	}
}

func (op *LoadTabOperation) NoWait() (bool, bool) {
	if op.input.NoWait != nil {
		return true, op.input.NoWaitValue
	} else {
		return false, false
	}
}

func (op *LoadTabOperation) SetTabId(tabId int) *LoadTabOperation {
	op.doEnsureNotStarted()
	op.input.TabId = tabId
	return op
}

func (op *LoadTabOperation) SetUrl(url string) *LoadTabOperation {
	op.doEnsureNotStarted()
	op.input.Url = url
	return op
}

func (op *LoadTabOperation) SetReplace(specified bool, replace bool) *LoadTabOperation {
	op.doEnsureNotStarted()

	if specified {
		op.input.ReplaceValue = replace
		op.input.Replace = &op.input.ReplaceValue
		return op
	} else {
		op.input.Replace = nil
		return op
	}
}

func (op *LoadTabOperation) SetNoWait(specified bool, noWait bool) *LoadTabOperation {
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

func (op *LoadTabOperation) Start(transport *Transport, callback func(op *LoadTabOperation)) {
	op.doStart(transport, protocol.LoadTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *LoadTabOperation) StartChannel(transport *Transport, channel chan *LoadTabOperation) {
	op.doStart(transport, protocol.LoadTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *LoadTabOperation) Execute(transport *Transport) (*protocol.Tab, error) {
	if err := op.doExecute(transport, protocol.LoadTabMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *LoadTabOperation) Result() (*protocol.Tab, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute tabs.reload
// method on a mindctrl web extension instance.
//
type ReloadTabOperation struct {
	GenericOperation
	input  protocol.ReloadTabInput
	output protocol.ReloadTabOutput
}

func ReloadTab(tabId int) *ReloadTabOperation {
	op := &ReloadTabOperation{}
	op.input.TabId = tabId
	op.input.BypassCache = nil
	op.input.NoWait = nil
	return op
}

func (op *ReloadTabOperation) TabId() int {
	return op.input.TabId
}

func (op *ReloadTabOperation) BypassCache() (bool, bool) {
	if op.input.BypassCache != nil {
		return true, op.input.BypassCacheValue
	} else {
		return false, false
	}
}

func (op *ReloadTabOperation) NoWait() (bool, bool) {
	if op.input.NoWait != nil {
		return true, op.input.NoWaitValue
	} else {
		return false, false
	}
}

func (op *ReloadTabOperation) SetTabId(tabId int) *ReloadTabOperation {
	op.doEnsureNotStarted()
	op.input.TabId = tabId
	return op
}

func (op *ReloadTabOperation) SetBypassCache(specified bool, bypassCache bool) *ReloadTabOperation {
	op.doEnsureNotStarted()

	if specified {
		op.input.BypassCacheValue = bypassCache
		op.input.BypassCache = &op.input.BypassCacheValue
		return op
	} else {
		op.input.BypassCache = nil
		return op
	}
}

func (op *ReloadTabOperation) SetNoWait(specified bool, noWait bool) *ReloadTabOperation {
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

func (op *ReloadTabOperation) Start(transport *Transport, callback func(op *ReloadTabOperation)) {
	op.doStart(transport, protocol.ReloadTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *ReloadTabOperation) StartChannel(transport *Transport, channel chan *ReloadTabOperation) {
	op.doStart(transport, protocol.ReloadTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *ReloadTabOperation) Execute(transport *Transport) (*protocol.Tab, error) {
	if err := op.doExecute(transport, protocol.ReloadTabMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *ReloadTabOperation) Result() (*protocol.Tab, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute tabs.activate
// method on a mindctrl web extension instance.
//
type ActivateTabOperation struct {
	GenericOperation
	input  protocol.ActivateTabInput
	output protocol.ActivateTabOutput
}

func ActivateTab(tabId int) *ActivateTabOperation {
	op := &ActivateTabOperation{}
	op.input.TabId = tabId
	return op
}

func (op *ActivateTabOperation) TabId() int {
	return op.input.TabId
}

func (op *ActivateTabOperation) SetTabId(tabId int) *ActivateTabOperation {
	op.doEnsureNotStarted()
	op.input.TabId = tabId
	return op
}

func (op *ActivateTabOperation) Start(transport *Transport, callback func(op *ActivateTabOperation)) {
	op.doStart(transport, protocol.ActivateTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *ActivateTabOperation) StartChannel(transport *Transport, channel chan *ActivateTabOperation) {
	op.doStart(transport, protocol.ActivateTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *ActivateTabOperation) Execute(transport *Transport) (*protocol.Tab, error) {
	if err := op.doExecute(transport, protocol.ActivateTabMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *ActivateTabOperation) Result() (*protocol.Tab, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute tabs.deactivate
// method on a mindctrl web extension instance.
//
type DeactivateTabOperation struct {
	GenericOperation
	input  protocol.DeactivateTabInput
	output protocol.DeactivateTabOutput
}

func DeactivateTab(tabId int) *DeactivateTabOperation {
	op := &DeactivateTabOperation{}
	op.input.TabId = tabId
	return op
}

func (op *DeactivateTabOperation) TabId() int {
	return op.input.TabId
}

func (op *DeactivateTabOperation) SetTabId(tabId int) *DeactivateTabOperation {
	op.doEnsureNotStarted()
	op.input.TabId = tabId
	return op
}

func (op *DeactivateTabOperation) Start(transport *Transport, callback func(op *DeactivateTabOperation)) {
	op.doStart(transport, protocol.DeactivateTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *DeactivateTabOperation) StartChannel(transport *Transport, channel chan *DeactivateTabOperation) {
	op.doStart(transport, protocol.DeactivateTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *DeactivateTabOperation) Execute(transport *Transport) (*protocol.Tab, error) {
	if err := op.doExecute(transport, protocol.DeactivateTabMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *DeactivateTabOperation) Result() (*protocol.Tab, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute tabs.mute
// method on a mindctrl web extension instance.
//
type MuteTabOperation struct {
	GenericOperation
	input  protocol.MuteTabInput
	output protocol.MuteTabOutput
}

func MuteTab(tabId int) *MuteTabOperation {
	op := &MuteTabOperation{}
	op.input.TabId = tabId
	return op
}

func (op *MuteTabOperation) TabId() int {
	return op.input.TabId
}

func (op *MuteTabOperation) SetTabId(tabId int) *MuteTabOperation {
	op.doEnsureNotStarted()
	op.input.TabId = tabId
	return op
}

func (op *MuteTabOperation) Start(transport *Transport, callback func(op *MuteTabOperation)) {
	op.doStart(transport, protocol.MuteTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *MuteTabOperation) StartChannel(transport *Transport, channel chan *MuteTabOperation) {
	op.doStart(transport, protocol.MuteTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *MuteTabOperation) Execute(transport *Transport) (*protocol.Tab, error) {
	if err := op.doExecute(transport, protocol.MuteTabMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *MuteTabOperation) Result() (*protocol.Tab, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute tabs.unmute
// method on a mindctrl web extension instance.
//
type UnmuteTabOperation struct {
	GenericOperation
	input  protocol.UnmuteTabInput
	output protocol.UnmuteTabOutput
}

func UnmuteTab(tabId int) *UnmuteTabOperation {
	op := &UnmuteTabOperation{}
	op.input.TabId = tabId
	return op
}

func (op *UnmuteTabOperation) TabId() int {
	return op.input.TabId
}

func (op *UnmuteTabOperation) SetTabId(tabId int) *UnmuteTabOperation {
	op.doEnsureNotStarted()
	op.input.TabId = tabId
	return op
}

func (op *UnmuteTabOperation) Start(transport *Transport, callback func(op *UnmuteTabOperation)) {
	op.doStart(transport, protocol.UnmuteTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *UnmuteTabOperation) StartChannel(transport *Transport, channel chan *UnmuteTabOperation) {
	op.doStart(transport, protocol.UnmuteTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *UnmuteTabOperation) Execute(transport *Transport) (*protocol.Tab, error) {
	if err := op.doExecute(transport, protocol.UnmuteTabMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *UnmuteTabOperation) Result() (*protocol.Tab, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute tabs.pin
// method on a mindctrl web extension instance.
//
type PinTabOperation struct {
	GenericOperation
	input  protocol.PinTabInput
	output protocol.PinTabOutput
}

func PinTab(tabId int) *PinTabOperation {
	op := &PinTabOperation{}
	op.input.TabId = tabId
	return op
}

func (op *PinTabOperation) TabId() int {
	return op.input.TabId
}

func (op *PinTabOperation) SetTabId(tabId int) *PinTabOperation {
	op.doEnsureNotStarted()
	op.input.TabId = tabId
	return op
}

func (op *PinTabOperation) Start(transport *Transport, callback func(op *PinTabOperation)) {
	op.doStart(transport, protocol.PinTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *PinTabOperation) StartChannel(transport *Transport, channel chan *PinTabOperation) {
	op.doStart(transport, protocol.PinTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *PinTabOperation) Execute(transport *Transport) (*protocol.Tab, error) {
	if err := op.doExecute(transport, protocol.PinTabMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *PinTabOperation) Result() (*protocol.Tab, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute tabs.unpin
// method on a mindctrl web extension instance.
//
type UnpinTabOperation struct {
	GenericOperation
	input  protocol.UnpinTabInput
	output protocol.UnpinTabOutput
}

func UnpinTab(tabId int) *UnpinTabOperation {
	op := &UnpinTabOperation{}
	op.input.TabId = tabId
	return op
}

func (op *UnpinTabOperation) TabId() int {
	return op.input.TabId
}

func (op *UnpinTabOperation) SetTabId(tabId int) *UnpinTabOperation {
	op.doEnsureNotStarted()
	op.input.TabId = tabId
	return op
}

func (op *UnpinTabOperation) Start(transport *Transport, callback func(op *UnpinTabOperation)) {
	op.doStart(transport, protocol.UnpinTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *UnpinTabOperation) StartChannel(transport *Transport, channel chan *UnpinTabOperation) {
	op.doStart(transport, protocol.UnpinTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *UnpinTabOperation) Execute(transport *Transport) (*protocol.Tab, error) {
	if err := op.doExecute(transport, protocol.UnpinTabMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *UnpinTabOperation) Result() (*protocol.Tab, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute tabs.move
// method on a mindctrl web extension instance.
//
type MoveTabOperation struct {
	GenericOperation
	input  protocol.MoveTabInput
	output protocol.MoveTabOutput
}

func MoveTab(tabId int, index int) *MoveTabOperation {
	op := &MoveTabOperation{}
	op.input.TabId = tabId
	op.input.Index = index
	return op
}

func (op *MoveTabOperation) TabId() int {
	return op.input.TabId
}

func (op *MoveTabOperation) Index() int {
	return op.input.Index
}

func (op *MoveTabOperation) WindowId() (bool, int) {
	if op.input.WindowId != nil {
		return true, op.input.WindowIdValue
	} else {
		return false, -2
	}
}

func (op *MoveTabOperation) SetTabId(tabId int) *MoveTabOperation {
	op.doEnsureNotStarted()
	op.input.TabId = tabId
	return op
}

func (op *MoveTabOperation) SetIndex(index int) *MoveTabOperation {
	op.doEnsureNotStarted()
	op.input.Index = index
	return op
}

func (op *MoveTabOperation) SetWindowId(specified bool, windowId int) *MoveTabOperation {
	op.doEnsureNotStarted()

	if specified {
		op.input.WindowIdValue = windowId
		op.input.WindowId = &op.input.WindowIdValue
		return op
	} else {
		op.input.WindowId = nil
		return op
	}
}

func (op *MoveTabOperation) Start(transport *Transport, callback func(op *MoveTabOperation)) {
	op.doStart(transport, protocol.MoveTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *MoveTabOperation) StartChannel(transport *Transport, channel chan *MoveTabOperation) {
	op.doStart(transport, protocol.MoveTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *MoveTabOperation) Execute(transport *Transport) (*protocol.Tab, error) {
	if err := op.doExecute(transport, protocol.MoveTabMethod, &op.input, &op.output); err != nil {
		return nil, err
	} else if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

func (op *MoveTabOperation) Result() (*protocol.Tab, error) {
	op.doEnsureFinished()

	if op.output.Success == false {
		return nil, errors.New(op.output.Message)
	} else {
		return &op.output.Result, nil
	}
}

// This operation provides a fluent interface to execute tabs.discard
// method on a mindctrl web extension instance.
//
type DiscardTabOperation struct {
	GenericOperation
	input  protocol.DiscardTabInput
	output protocol.DiscardTabOutput
}

func DiscardTab(tabId int) *DiscardTabOperation {
	op := &DiscardTabOperation{}
	op.input.TabId = tabId
	return op
}

func (op *DiscardTabOperation) TabId() int {
	return op.input.TabId
}

func (op *DiscardTabOperation) SetTabId(tabId int) *DiscardTabOperation {
	op.doEnsureNotStarted()
	op.input.TabId = tabId
	return op
}

func (op *DiscardTabOperation) Start(transport *Transport, callback func(op *DiscardTabOperation)) {
	op.doStart(transport, protocol.DiscardTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *DiscardTabOperation) StartChannel(transport *Transport, channel chan *DiscardTabOperation) {
	op.doStart(transport, protocol.DiscardTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *DiscardTabOperation) Execute(transport *Transport) error {
	if err := op.doExecute(transport, protocol.DiscardTabMethod, &op.input, &op.output); err != nil {
		return err
	} else if op.output.Success == false {
		return errors.New(op.output.Message)
	} else {
		return nil
	}
}

func (op *DiscardTabOperation) Result() error {
	op.doEnsureFinished()

	if op.output.Success == false {
		return errors.New(op.output.Message)
	} else {
		return nil
	}
}

// This operation provides a fluent interface to execute tabs.remove
// method on a mindctrl web extension instance.
//
type RemoveTabOperation struct {
	GenericOperation
	input  protocol.RemoveTabInput
	output protocol.RemoveTabOutput
}

func RemoveTab(tabId int) *RemoveTabOperation {
	op := &RemoveTabOperation{}
	op.input.TabId = tabId
	return op
}

func (op *RemoveTabOperation) TabId() int {
	return op.input.TabId
}

func (op *RemoveTabOperation) SetTabId(tabId int) *RemoveTabOperation {
	op.doEnsureNotStarted()
	op.input.TabId = tabId
	return op
}

func (op *RemoveTabOperation) Start(transport *Transport, callback func(op *RemoveTabOperation)) {
	op.doStart(transport, protocol.RemoveTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		callback(op)
	})
}

func (op *RemoveTabOperation) StartChannel(transport *Transport, channel chan *RemoveTabOperation) {
	op.doStart(transport, protocol.RemoveTabMethod, &op.input, &op.output, func(m string, a, r interface{}, err error) {
		op.doFinish(err)
		channel <- op
	})
}

func (op *RemoveTabOperation) Execute(transport *Transport) error {
	if err := op.doExecute(transport, protocol.RemoveTabMethod, &op.input, &op.output); err != nil {
		return err
	} else if op.output.Success == false {
		return errors.New(op.output.Message)
	} else {
		return nil
	}
}

func (op *RemoveTabOperation) Result() error {
	op.doEnsureFinished()

	if op.output.Success == false {
		return errors.New(op.output.Message)
	} else {
		return nil
	}
}
