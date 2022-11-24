package reader

import (
	"errors"
	"os"
)

var NoEntityErr error = errors.New("No upcomming entities")

// home directory
var home = func() string {
	res, _ := os.UserHomeDir()
	return res
}()

// Default values for config type
var (
	EmacsDirDef           = home + "/.emacs.d"
	ScanIntDef            = 300
	BeforeNotificationDef = 30
	NotifyCommandDef      = "notify-send"
)

// configuration locations
var (
	ConfigDir       = home + "/.config/agenda-notification"
	LogFile         = ConfigDir + "/log.log"
	ConfigFileLoc   = ConfigDir + "/config.yaml"
	ExportScriptLoc = ConfigDir + "/exportScript"
	IconLoc         = ConfigDir + "/icon.png"
)

// PossibleEmacsConfigLocations contains directories to check for emacs init files
var PossibleEmacsConfigLocations = []string{
	home + "/.emacs",
	home + "/.emacs.el",
	home + "/.emacs.d/init.el",
	home + "/.config/emacs/init.el",
}

// PossibleDoomConfigLocations contains possible locations for doom emacs config
var PossibleDoomConfigLocations = []string{
	home + "/.doom.d/",
	home + "/.config/.doom.d",
}

// Agenda outputs flags, this helps to avoid reading warnings from emacs
const (
	AgendaStart = "\"STARTAGENDA\""
	AgendaEnd   = "\"ENDAGENDA\""
)

// DoomExporter is a doomscript file which exports agenda informations in case of doom emacs
const DoomExporter = `#!/usr/bin/env doomscript
(require 'doom-start)
(let ((inhibit-message t))
(print "STARTAGENDA")
(org-batch-agenda-csv "a")
(print "ENDAGENDA"))
`

const EmacsExporter = `
(let ((inhibit-message t))
(message "Listen to me, you!")
(load-file "/home/ghd/.emacs.d/init.el")
(print "STARTAGENDA")
(org-batch-agenda-csv "a")
(print "ENDAGENDA"))`

const TimeFormat = "2006-01-02_15:04"
