package protocol

// Names of the RPC service methods recognized by the server. Each
// method will have the corresponding input and output structures
// for the args and reply of the RPC call.
//
const (
	QueryDocumentMethod = "documents.query"

	FindDownloadsMethod  = "downloads.find"
	GetDownloadMethod    = "downloads.get"
	CreateDownloadMethod = "downloads.create"
	PauseDownloadMethod  = "downloads.pause"
	ResumeDownloadMethod = "downloads.resume"
	CancelDownloadMethod = "downloads.cancel"
	RemoveDownloadMethod = "downloads.remove"

	GetBrowserInfoMethod  = "info.get_browser"
	GetPlatformInfoMethod = "info.get_platform"

	PingMethod = "ping"

	FindTabsMethod      = "tabs.find"
	GetTabMethod        = "tabs.get"
	GetCurrentTabMethod = "tabs.get_current"
	CreateTabMethod     = "tabs.create"
	LoadTabMethod       = "tabs.load"
	ReloadTabMethod     = "tabs.reload"
	ActivateTabMethod   = "tabs.activate"
	DeactivateTabMethod = "tabs.deactivate"
	MuteTabMethod       = "tabs.mute"
	UnmuteTabMethod     = "tabs.unmute"
	PinTabMethod        = "tabs.pin"
	UnpinTabMethod      = "tabs.unpin"
	MoveTabMethod       = "tabs.move"
	DiscardTabMethod    = "tabs.discard"
	RemoveTabMethod     = "tabs.remove"

	FindWindowsMethod      = "windows.find"
	GetWindowMethod        = "windows.get"
	GetCurrentWindowMethod = "windows.get_current"
	CreateWindowMethod     = "windows.create"
	MoveWindowMethod       = "windows.move"
	ResizeWindowMethod     = "windows.resize"
	MinimizeWindowMethod   = "windows.minimize"
	MaximizeWindowMethod   = "windows.maximize"
	FullscreenWindowMethod = "windows.fullscreen"
	RestoreWindowMethod    = "windows.restore"
	FocusWindowMethod      = "windows.focus"
	UnfocusWindowMethod    = "windows.unfocus"
	RemoveWindowMethod     = "windows.remove"
)

// Details of the browser. It contains information about the
// browser running on the other side.
//
type BrowserInfo struct {
	Name    string `json:"name"`    // name of the browser
	Version string `json:"version"` // version of the browser
}

// Details of the platform. It contains information about the
// platform the browser at the other side is running on.
//
// The structure is adapted from the runtime.PlatformInfo type of
// the Web Extension API. The details on the structure can be
// found in:
//
//   - https://developer.chrome.com/docs/extensions/reference/runtime/#type-PlatformInfo
//   - https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/runtime/PlatformInfo
//
type PlatformInfo struct {
	Os   string `json:"os"`   // name of the operating system the browser is running on
	Arch string `json:"arch"` // architecture of the processor the browser is running on
}

// Details of a single download.
//
// The structure is adapted from the downloads.DownloadItem type of
// the Web Extension API. The details on the structure can be found
// in:
//
//   - https://developer.chrome.com/docs/extensions/reference/downloads/#type-DownloadItem
//   - https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/downloads/DownloadItem
//
type Download struct {
	Id            int    `json:"id"`            // id of the download
	Url           string `json:"url"`           // url of the download
	Filename      string `json:"filename"`      // name of the output file
	Referrer      string `json:"referrer"`      // referrer of the download
	Mime          string `json:"mime"`          // media type of the download
	State         string `json:"state"`         // state of the download
	Paused        bool   `json:"paused"`        // whether the download is paused
	CanResume     bool   `json:"canResume"`     // whether the download is eligible for resume
	Error         string `json:"error"`         // error occured during the download
	StartTime     string `json:"startTime"`     // time when the download starts
	Filesize      int64  `json:"fileSize"`      // size of the output file
	TotalBytes    int64  `json:"totalBytes"`    // number of bytes to be downloaded
	ReceivedBytes int64  `json:"bytesReceived"` // number of bytes received thus far
}

// Details for a single tab.
//
// The structure is adapted from the tabs.Tab type of the Web
// Extension API. The details on the structure can be found in:
//
//   - https://developer.chrome.com/docs/extensions/reference/tabs/#type-Tab
//   - https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/tabs/Tab
//
type Tab struct {
	Id          int       `json:"id"`              // id of the tab
	WindowId    int       `json:"windowId"`        // window the tab belongs to
	Index       int       `json:"index"`           // index of the tab in the window
	Width       int       `json:"width"`           // width of the tab
	Height      int       `json:"height"`          // height of the tab
	Url         string    `json:"url"`             // url of the document in the tab
	Title       string    `json:"title"`           // title of the document in the tab
	FavIcon     string    `json:"favIconUrl"`      // url of favicon of the document in the tab
	Status      string    `json:"status"`          // loading Status of the document in the tab ("loading" or "completed")
	Active      bool      `json:"active"`          // whether the tab is active or not
	Highlighted bool      `json:"highlighted"`     // whether the tab is highlighted or not
	Pinned      bool      `json:"pinned"`          // whether the tab is pinned or not
	Hidden      bool      `json:"hidden"`          // whether the tab is hidden or not
	Discarded   bool      `json:"discarded"`       // whether the tab is discarded or not
	Discardable bool      `json:"autoDiscardable"` // whether the tab can be discarded or not
	Attention   bool      `json:"attention"`       // whether the tab requires attention or not
	Audible     bool      `json:"audible"`         // whether the tab is making sound or not
	Muted       MutedInfo `json:"mutedInfo"`       // mute status of the tab
}

// Mute status for a tab. It indicates if the given tab is muted,
// and if the tab is muted, the reason behind it.
//
// The structure is adapted from the tabs.MutedInfo type of the Web
// Extension API. The details on the structure can be found in:
//
//   - https://developer.chrome.com/docs/extensions/reference/tabs/#type-MutedInfo
//   - https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/tabs/MutedInfo
//
type MutedInfo struct {
	Muted  bool   `json:"muted"`  // whether the tab is muted or not
	Reason string `json:"reason"` // explanation on why the tab is muted ("user", "capture" or "extension")
}

// Details for a single window.
//
// The structure is adapted from the tabs.MutedInfo type of the Web
// Extension API. The details on the structure can be found in:
//
//   - https://developer.chrome.com/docs/extensions/reference/windows/#type-Window
//   - https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/windows/Window
//
type Window struct {
	Id          int    `json:"id"`          // id of the window
	Type        string `json:"type"`        // type of the window ("normal", "popup", "panel" or "devtools")
	Width       int    `json:"width"`       // width of the window
	Left        int    `json:"left"`        // horizontal position of the window
	Top         int    `json:"top"`         // vertical position of the window
	Height      int    `json:"height"`      // height of the window
	Title       string `json:"title"`       // title of the window
	State       string `json:"state"`       // state of the window ("minimized", "maximized", "fullscreen" or "docked")
	Focused     bool   `json:"focused"`     // whether the window is focused or not
	AlwaysOnTop bool   `json:"alwaysOnTop"` // whether the window is always on top or not
	Tabs        []Tab  `json:"tabs"`        // list of tabs in the window
}

// Common fields of method outputs. The structure can be embedded by
// the output structs to include these fields automatically without
// repeating them everywhere.
//
type GenericOutput struct {
	Success  bool   `json:"success"`  // whether the operation is successful
	Category string `json:"category"` // category of the error; exists only when success is false
	Message  string `json:"message"`  // explanation of the error; exists only when success is false
}

// Input for document.query RPC method. Currently the method
// requires ID of the target tab and the GraphQL query to be
// executed. The method also optionally accepts the name of the
// operation and any variables required by the GraphQL query.
//
type QueryDocumentInput struct {
	TabId     int                    `json:"tabId"`
	Query     string                 `json:"query"`
	Operation string                 `json:"operation,omitempty"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

// Output for document.query PRC method. Besides the usual fields,
// the output also contains the result of the query which would be
// populated if the method is completed successfully. Note that
// the result is stored as a raw JSON message so that caller can
// do their own unmarshalling.
//
type QueryDocumentOutput struct {
	GenericOutput
	Result interface{} `json:"result"`
}

// Input for downloads.find RPC method. The method optionally
// accepts an URL match pattern and download state to filter
// out any downloads that does not match the given criteria.
//
type FindDownloadsInput struct {
	Url        *string `json:"url,omitempty"`
	State      *string `json:"state,omitempty"`
	UrlValue   string  `json:"-"`
	StateValue string  `json:"-"`
}

// Output for downloads.find PRC method. Besides the usual fields,
// the output also contains the list of matching downloads which
// would be populated if the method is completed successfully.
//
type FindDownloadsOutput struct {
	GenericOutput
	Result []Download `json:"result"`
}

// Input for downloads.get RPC method. The method requires the
// ID of the target download.
//
type GetDownloadInput struct {
	DownloadId int `json:"downloadId"`
}

// Output for downloads.get RPC method. Besides the usual fields,
// the output also contains details of the target download which
// would be populated if the method is completed successfully.
//
type GetDownloadOutput struct {
	GenericOutput
	Result Download `json:"result"`
}

// Input for downloads.create RPC method. The method requires
// the URL of the target resource and name of the file where the
// downloaded data is saved. The method also optionally accepts
// the referrer for the download request, and whether the method
// should wait for the download to stop before returning.
//
type CreateDownloadInput struct {
	Url           string  `json:"url"`
	Filename      string  `json:"filename"`
	Referrer      *string `json:"referrer,omitempty"`
	NoWait        *bool   `json:"noWait,omitempty"`
	ReferrerValue string  `json:"-"`
	NoWaitValue   bool    `json:"-"`
}

// Output for downloads.create RPC method. Besides the usual fields,
// the output also contains details of the new download which would
// be populated if the method is completed successfully.
//
type CreateDownloadOutput struct {
	GenericOutput
	Result Download `json:"result"`
}

// Input for downloads.pause RPC method. The method requires the ID
// of the target download.
//
type PauseDownloadInput struct {
	DownloadId int `json:"downloadId"`
}

// Output for downloads.pause RPC method. Besides the usual fields,
// the output also contains details of the paused download which
// would be populated if the method is completed successfully.
//
type PauseDownloadOutput struct {
	GenericOutput
	Result Download `json:"result"`
}

// Input for downloads.resume RPC method. The method requires the
// ID of the target download. The method also optionally accepts
// whether the method should wait for the download to stop before
// returning.
//
type ResumeDownloadInput struct {
	DownloadId  int   `json:"downloadId"`
	NoWait      *bool `json:"noWait,omitempty"`
	NoWaitValue bool  `json:"-"`
}

// Output for downloads.resume RPC method. Besides the usual fields,
// the output also contains details of the resumed download which
// would be populated if the method is completed successfully.
//
type ResumeDownloadOutput struct {
	GenericOutput
	Result Download `json:"result"`
}

// Input for downloads.cancel RPC method. The method requires the
// ID of the target download.
//
type CancelDownloadInput struct {
	DownloadId int `json:"downloadId"`
}

// Output for downloads.cancel RPC method. Besides the usual fields,
// the output also contains details of the cancelled download which
// would be populated if the method is completed successfully.
//
type CancelDownloadOutput struct {
	GenericOutput
	Result Download `json:"result"`
}

// Input for downloads.remove RPC method. The method requires the
// ID of the target download.
//
type RemoveDownloadInput struct {
	DownloadId int `json:"downloadId"`
}

// Output for downloads.remove RPC method. The output does not
// contain any extras besides the usual fields.
//
type RemoveDownloadOutput struct {
	GenericOutput
}

// Input for info.get_browser RPC method. The method does not
// require any extra data.
//
type GetBrowserInfoInput struct {
	// empty
}

// Output for info.get_browser PRC method. Besides the usual
// fields, the output also contains details of the browser which
// would be populated if the method is completed successfully.
//
type GetBrowserInfoOutput struct {
	GenericOutput
	Result BrowserInfo `json:"result"`
}

// Input for info.get_platform RPC method. The method does not
// require any extra data.
//
type GetPlatformInfoInput struct {
	// empty
}

// Output for info.get_platform PRC method. Besides the usual
// fields, the output also contains details of the platform which
// would be populated if the method is completed successfully.
//
type GetPlatformInfoOutput struct {
	GenericOutput
	Result PlatformInfo `json:"result"`
}

// Input for ping RPC method. The method does not require any extra
// data.
//
type PingInput struct {
	// empty
}

// Output for the ping RPC method. The output does not contain any
// extras besides the usual fields.
//
type PingOutput struct {
	GenericOutput
}

// Input for tabs.find RPC method. The method optionally accepts
// window ID, url match pattern, status, etc to filter out any tabs
// that do not match the given criteria.
//
type FindTabsInput struct {
	WindowId       *int    `json:"windowId,omitempty"`
	Url            *string `json:"url,omitempty"`
	Status         *string `json:"status,omitempty"`
	Active         *bool   `json:"active,omitempty"`
	Audible        *bool   `json:"audible,omitempty"`
	Discarded      *bool   `json:"discarded,omitempty"`
	Muted          *bool   `json:"muted,omitempty"`
	Pinned         *bool   `json:"pinned,omitempty"`
	WindowIdValue  int     `json:"-"`
	UrlValue       string  `json:"-"`
	StatusValue    string  `json:"-"`
	ActiveValue    bool    `json:"-"`
	AudibleValue   bool    `json:"-"`
	DiscardedValue bool    `json:"-"`
	MutedValue     bool    `json:"-"`
	PinnedValue    bool    `json:"-"`
}

// Output for tabs.find PRC method. Besides the usual fields, the
// output also contains the list of matching tabs which would be
// populated if the method is completed successfully.
//
type FindTabsOutput struct {
	GenericOutput
	Result []Tab `json:"result"`
}

// Input for tabs.get RPC method. The method requires the ID of the
// target tab.
//
type GetTabInput struct {
	TabId int `json:"tabId"`
}

// Output for tabs.get PRC method. Besides the usual fields, the
// output also contains details of the target tabs which would be
// populated if the method is completed successfully.
//
type GetTabOutput struct {
	GenericOutput
	Result Tab `json:"result"`
}

// Input for tabs.get_current RPC method. The argument does not
// require any data.
//
type GetCurrentTabInput struct {
	// empty
}

// Output for tabs.get_current PRC method. Besides the usual fields,
// the output also contains details of the current tabs which would
// be populated if the method is completed successfully.
//
type GetCurrentTabOutput struct {
	GenericOutput
	Result Tab `json:"result"`
}

// Input for tabs.create RPC method. The method optionally accepts
// the window the tab will be created in, whether the tab should
// be activated on creation, URL to be loaded in the tab, and
// whether the method should wait for the tab to finish loading
// before returning.
//
type CreateTabInput struct {
	WindowId      *int    `json:"windowId,omitempty"`
	Url           *string `json:"url,omitempty"`
	Active        *bool   `json:"active,omitempty"`
	NoWait        *bool   `json:"noWait,omitempty"`
	WindowIdValue int     `json:"-"`
	UrlValue      string  `json:"-"`
	ActiveValue   bool    `json:"-"`
	NoWaitValue   bool    `json:"-"`
}

// Output for tabs.create RPC method. Besides the usual fields, the
// output also contains details of the created tabs which would be
// populated if the method is completed successfully.
//
type CreateTabOutput struct {
	GenericOutput
	Result Tab `json:"result"`
}

// Input for tabs.load RPC method. The method requires the ID of
// the target tab and the URL to be loaded in the tab. The method
// optionally accepts whether the incoming page should replace the
// current page in the tab's history stack, and whether the method
// should wait for the tab to finish loading before returning.
//
type LoadTabInput struct {
	TabId        int    `json:"tabId"`
	Url          string `json:"url"`
	Replace      *bool  `json:"replace,omitempty"`
	NoWait       *bool  `json:"noWait,omitempty"`
	ReplaceValue bool   `json:"-"`
	NoWaitValue  bool   `json:"-"`
}

// Output for tabs.load RPC method. Besides the usual fields, the
// output also contains details of the loaded tabs which would be
// populated if the method is completed successfully.
//
type LoadTabOutput struct {
	GenericOutput
	Result Tab `json:"result"`
}

// Input for tabs.reload RPC method. The method requires the ID of
// the target tab. The method optionally accepts whether the page
// cache is bypassed, and whether the method should wait for the
// tab to finish loading before returning.
//
type ReloadTabInput struct {
	TabId            int   `json:"tabId"`
	BypassCache      *bool `json:"bypassCache,omitempty"`
	NoWait           *bool `json:"noWait,omitempty"`
	BypassCacheValue bool  `json:"-"`
	NoWaitValue      bool  `json:"-"`
}

// Output for tabs.reload RPC method. Besides the usual fields, the
// output also contains details of the reloaded tabs which would be
// populated if the method is completed successfully.
//
type ReloadTabOutput struct {
	GenericOutput
	Result Tab `json:"result"`
}

// Input for tabs.activate RPC method. The method requires the ID
// of the target tab.
//
type ActivateTabInput struct {
	TabId int `json:"tabId"`
}

// Output for tabs.activate RPC method. Besides the usual fields,
// the output also contains details of the activated tabs which
// would be populated if the method is completed successfully.
//
type ActivateTabOutput struct {
	GenericOutput
	Result Tab `json:"result"`
}

// Input for tabs.deactivate RPC method. The method requires the ID
// of the target tab.
//
type DeactivateTabInput struct {
	TabId int `json:"tabId"`
}

// Output for tabs.deactivate RPC method. Besides the usual fields,
// the output also contains details of the deactivated tabs which
// would be populated if the method is completed successfully.
//
type DeactivateTabOutput struct {
	GenericOutput
	Result Tab `json:"result"`
}

// Input for tabs.mute RPC method. The method requires the ID of
// the target tab.
//
type MuteTabInput struct {
	TabId int `json:"tabId"`
}

// Output for tabs.mute RPC method. Besides the usual fields, the
// output also contains details of the muted tabs which would be
// populated if the method is completed successfully.
//
type MuteTabOutput struct {
	GenericOutput
	Result Tab `json:"result"`
}

// Input for tabs.unmute RPC method. The method requires the ID of
// the target tab.
//
type UnmuteTabInput struct {
	TabId int `json:"tabId"`
}

// Output for tabs.unmute RPC method. Besides the usual fields, the
// output also contains details of the unmuted tabs which would be
// populated if the method is completed successfully.
//
type UnmuteTabOutput struct {
	GenericOutput
	Result Tab `json:"result"`
}

// Input for tabs.pin RPC method. The method requires the ID of
// the target tab.
//
type PinTabInput struct {
	TabId int `json:"tabId"`
}

// Output for tabs.pin RPC method. Besides the usual fields, the
// output also contains details of the pinned tabs which would be
// populated if the method is completed successfully.
//
type PinTabOutput struct {
	GenericOutput
	Result Tab `json:"result"`
}

// Input for tabs.unpin RPC method. The method requires the ID of
// the target tab.
//
type UnpinTabInput struct {
	TabId int `json:"tabId"`
}

// Output for tabs.unpin RPC method. Besides the usual fields, the
// output also contains details of the unpinned tabs which would be
// populated if the method is completed successfully.
//
type UnpinTabOutput struct {
	GenericOutput
	Result Tab `json:"result"`
}

// Input for tabs.move RPC method. The method requires the ID of
// the target tab and the index to be moved to. The method also
// optionally accepts the window where the tab should move to.
//
type MoveTabInput struct {
	TabId         int  `json:"tabId"`
	Index         int  `json:"index"`
	WindowId      *int `json:"windowId,omitempty"`
	WindowIdValue int  `json:"-"`
}

// Output for tabs.move RPC method. Besides the usual fields, the
// output also contains details of the moeed tabs which would be
// populated if the method is completed successfully.
//
type MoveTabOutput struct {
	GenericOutput
	Result Tab `json:"result"`
}

// Input for tabs.discard RPC method. The method requires the ID of
// the target tab.
//
type DiscardTabInput struct {
	TabId int `json:"tabId"`
}

// Output for tabs.discard RPC method. The output does not contain
// any extras besides the usual fields.
//
type DiscardTabOutput struct {
	GenericOutput
}

// Input for tabs.remove RPC method. The method requires the ID of
// the target tab.
//
type RemoveTabInput struct {
	TabId int `json:"tabId"`
}

// Output for tabs.remove RPC method. The output does not contain
// any extras besides the usual fields.
//
type RemoveTabOutput struct {
	GenericOutput
}

// Input for windows.find RPC method. The method does not require
// any extra data.
//
type FindWindowsInput struct {
	// empty
}

// Output for windows.find PRC method. Besides the usual fields, the
// output also contains the list of matching windows which would be
// populated if the method is completed successfully.
//
type FindWindowsOutput struct {
	GenericOutput
	Result []Window `json:"result"`
}

// Input for windows.get RPC method. The method requires the ID of
// the target window.
//
type GetWindowInput struct {
	WindowId int `json:"windowId"`
}

// Output for windows.get RPC method. Besides the usual fields, the
// output also contains details of the target windows which would
// be populated if the method is completed successfully.
//
type GetWindowOutput struct {
	GenericOutput
	Result Window `json:"result"`
}

// Input for windows.get_current RPC method. The method does not
// require any data.
//
type GetCurrentWindowInput struct {
}

// Output for windows.get_current RPC method. Besides the usual
// fields, the output also contains details of the current window
// which would be populated if the method is completed successfully.
//
type GetCurrentWindowOutput struct {
	GenericOutput
	Result Window `json:"result"`
}

// Input for windows.create RPC method. Currently the method
// requires no data, but more will be added later.
//
type CreateWindowInput struct {
	// TODO: extra data
}

// Output for windowss.create RPC method. Besides the usual fields,
// the output also contains details of the created window which
// would be populated if the method is completed successfully.
//
type CreateWindowOutput struct {
	GenericOutput
	Result Window `json:"result"`
}

// Input for windows.move RPC method. The method requires the ID
// of the target window and the new position of the window.
//
type MoveWindowInput struct {
	WindowId int `json:"windowId"`
	Left     int `json:"left"`
	Top      int `json:"top"`
}

// Output for windows.move RPC method. Besides the usual fields,
// the output also contains details of the moved window which
// would be populated if the method is completed successfully.
//
type MoveWindowOutput struct {
	GenericOutput
	Result Window `json:"result"`
}

// Input for windows.resize RPC method. The method requires the
// ID of the target window and the new dimension of the window.
//
type ResizeWindowInput struct {
	WindowId int `json:"windowId"`
	Width    int `json:"width"`
	Height   int `json:"height"`
}

// Output for windows.resize RPC method. Besides the usual fields,
// the output also contains details of the resized window which
// would be populated if the method is completed successfully.
//
type ResizeWindowOutput struct {
	GenericOutput
	Result Window `json:"result"`
}

// Input for windows.minimize RPC method. The method requires the
// ID of the target window.
//
type MinimizeWindowInput struct {
	WindowId int `json:"windowId"`
}

// Output for windows.minimize RPC method. Besides the usual fields,
// the output also contains details of the minimized window which
// would be populated if the method is completed successfully.
//
type MinimizeWindowOutput struct {
	GenericOutput
	Result Window `json:"result"`
}

// Input for windows.maximize RPC method. The method requires the
// ID of the target window.
//
type MaximizeWindowInput struct {
	WindowId int `json:"windowId"`
}

// Output for windows.maximize RPC method. Besides the usual fields,
// the output also contains details of the maximized window which
// would be populated if the method is completed successfully.
//
type MaximizeWindowOutput struct {
	GenericOutput
	Result Window `json:"result"`
}

// Input for windows.fullscreen RPC method. The method requires the
// ID of the target window.
//
type FullscreenWindowInput struct {
	WindowId int `json:"windowId"`
}

// Output for windows.fullscreen RPC method. Besides the usual
// fields, the output also contains details of the fullscreened
// window which would be populated if the method is completed
// successfully.
//
type FullscreenWindowOutput struct {
	GenericOutput
	Result Window `json:"result"`
}

// Input for windows.restore RPC method. The method requires the
// ID of the target window.
//
type RestoreWindowInput struct {
	WindowId int `json:"windowId"`
}

// Output for windows.restore RPC method. Besides the usual fields,
// the output also contains details of the restored window which
// would be populated if the method is completed successfully.
//
type RestoreWindowOutput struct {
	GenericOutput
	Result Window `json:"result"`
}

// Input for windows.focus RPC method. The method requires the ID
// of the target window.
//
type FocusWindowInput struct {
	WindowId int `json:"windowId"`
}

// Output for windows.focus RPC method. Besides the usual fields,
// the output also contains details of the focused window which
// would be populated if the method is completed successfully.
//
type FocusWindowOutput struct {
	GenericOutput
	Result Window `json:"result"`
}

// Input for windows.unfocus RPC method. The method requires the
// ID of the target window.
//
type UnfocusWindowInput struct {
	WindowId int `json:"windowId"`
}

// Output for windows.unfocus RPC method. Besides the usual fields,
// the output also contains details of the unfocused window which
// would be populated if the method is completed successfully.
//
type UnfocusWindowOutput struct {
	GenericOutput
	Result Window `json:"result"`
}

// Input for windows.remove RPC method. The method requires the ID
// of the target window.
//
type RemoveWindowInput struct {
	WindowId int `json:"windowId"`
}

// Output for windows.remove RPC method. The output does not contain
// any extras besides the usual fields.
//
type RemoveWindowOutput struct {
	GenericOutput
}
