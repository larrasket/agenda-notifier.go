package main

import "errors"

// Default values for config type
const (
	EmacsDirDef           = "~/.emacs.d"
	ScanIntDef            = 300
	BeforeNotificationDef = 30
	NotifyCommandDef      = "notify-send"
)

// Other values
const (
	ConfigDirDef = "~/.config/agenda-notification"
	doomScript   = "doomscript"
)

// messages
const (
	doomFound      = "[Inf]Found doom emacs"
	doomNotFound   = "[Inf]Did not find doom emacs"
	emacsFound     = "[Inf]Found emacs at" // NEEDS PARAMS
	nonFound       = "[ERR]Couldn't find doom nor emacs, please define it at " + ConfigDirDef
	WriteConfError = "[ERR]Couldn't write config at " + ConfigDirDef
)

var PossibleEmacsConfigLocations = []string{
	"~/.emacs",
	"~/.emacs.el",
	"~/.emacs.d/init.el",
	"~/.config/emacs/init.el",
}

var PossibleDoomConfigLocations = []string{
	"~/.doom.d/",
	"~/.config/.doom.d",
}
var ErrNotFound = errors.New(nonFound)
