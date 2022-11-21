package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Config Type to generate a configuration file.
type Config struct {
	Doom               bool   `json:"doom"`               // whether user uses doom or not.
	EmacsLoc           string `json:"emacsDir"`           // emacs inti location                   (if !doom)
	ScanInt            int    `json:"scanInt"`            // scan interval.                        (in second)
	BeforeNotification int    `json:"beforeNotification"` // time before notifing                  (in second)
	NotifyCommand      string `json:"notifyCommand"`      // command to use for sending notification
	DoomScript         string `json:"doomScriptLoc"`      // doomscript binary location
}

// MakeConfig returns default configuration data
func MakeConfig(doom bool, init string) Config {
	conf := Config{
		doom,
		init,
		ScanIntDef,
		BeforeNotificationDef,
		NotifyCommandDef,
		EmacsDirDef + "/bin/doomscript",
	}
	return conf
}

// ReadConfig reads configuration file
func ReadConfig() (conf Config, err error) {
	fconf, err := os.ReadFile(ConfigFileLoc)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(fconf, &conf)
	return
}

func InitConfig() (doom bool, err error) {
	doom = false
	var loc *string
	doom, loc = IsDoom()
	if !doom {
		fmt.Println(doomFoundErr)
		_, loc = IsEmacs()
	}
	if loc != nil {
		fmt.Println(emacsFound + *loc)
	} else {
		fmt.Println(FoundErr)
		err = ErrNotFound
		return
	}
	conf, _ := yaml.Marshal(MakeConfig(doom, *loc))
	mod := os.FileMode(0777)
	err = os.Mkdir(ConfigDir, mod)
	err = os.WriteFile(ConfigFileLoc, conf, mod)
	err = os.WriteFile(ExportScriptLoc, makeScript(doom, *loc), mod)
	return
}

func makeScript(doom bool, loc string) []byte {
	if doom {
		return []byte(DoomExporter)
	} else {
		return []byte(fmt.Sprintf(EmacsExporter, loc))
	}
}
