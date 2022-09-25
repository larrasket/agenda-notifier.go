package main

import (
	"errors"
	"os"
)

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
	ConfigFileLoc   = ConfigDir + "/config.yaml"
	ExportScriptLoc = ConfigDir + "/exportScript"
)

// user cli messages
var (
	doomFoundErr = "[Inf]Did not find doom emacs"
	emacsFound   = "[Inf]Found emacs at "
	FoundErr     = "[ERR]Couldn't find doom nor emacs, please define it at " + ConfigDir // NEEDS PARAMS
	WriteConfErr = "[ERR]Couldn't write config at " + ConfigDir                          // NEEDS PARAMS
	ReadConfErr  = "[ERR]Couldn't read config file"
	ReadDataErr  = "[ERR]Couldn't reed agenda data"
	notFound     = "not found"
	noncomming   = "no upcoming entities"
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
	AgendaStart = "\"AGENDASTART\""
	AgendaEnd   = "\"ENDAGENDA\""
)

// DoomExporter is a doomscript file which exports agenda informations in case of doom emacs
const DoomExporter = `#!/usr/bin/env doomscript
(require 'doom-start)
(print "AGENDASTART")
(org-batch-agenda-csv "a")
(print "ENDAGENDA") `

const EmacsExporter = `(load-file "%s")
(print "STARTAGENDA")
(org-batch-agenda-csv "a")
(print "ENDAGENDA") `

var ErrNotFound = errors.New(notFound)
var ErrNoUpcomming = errors.New(noncomming)

const timeFormat = "15:04"

// TODO emacs script
