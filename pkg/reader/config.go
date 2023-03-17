package reader

import (
	"embed"
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Config Type to generate a configuration file.
type Config struct {
	Doom               bool   `yaml:"doom"`               // whether user uses doom or not.
	EmacsLoc           string `yaml:"emacsInitFile"`      // emacs init.el location                (if !doom)
	ScanInt            int    `yaml:"scanInt"`            // scan interval.                        (in second)
	BeforeNotification int    `yaml:"beforeNotification"` // time before notifing                  (in second)
	// NotifyCommand      string `yaml:"notifyCommand"`      // command to use for sending notification
	DoomScriptLoc string `yaml:"doomScriptLoc"` // doomscript binary location
}

// MakeConfig returns default configuration data
func MakeConfig(doom bool, init string) Config {
	conf := Config{
		doom,
		init,
		ScanIntDef,
		BeforeNotificationDef,
		// NotifyCommandDef,
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

func InitConfig() error {
	var loc *string
	var doom bool
	loc = IsDoom()

	if loc == nil {
		loc = IsEmacs()
	} else {
		doom = true
	}

	if loc == nil {
		return errors.New("Emacs was not found")
	}

	return writeConf(doom, loc)
}

//go:embed icon.png
var f embed.FS

func writeConf(doom bool, loc *string) error {
	conf, _ := yaml.Marshal(MakeConfig(doom, *loc))
	mod := os.FileMode(0777)
	err := os.MkdirAll(ConfigDir, mod)
	if err != nil {
		return err
	}
	err = os.WriteFile(ConfigFileLoc, conf, mod)
	if err != nil {
		return err
	}
	err = os.WriteFile(ExportScriptLoc, makeScript(doom, *loc), mod)
	img, err := f.ReadFile("icon.png")
	_ = os.WriteFile(IconLoc, img, mod)
	return err
}

func makeScript(doom bool, loc string) []byte {
	if doom {
		return []byte(DoomExporter)
	} else {
		return []byte(fmt.Sprintf(EmacsExporter, loc))
	}
}

// IsDoom checks if doom was found
func IsDoom() *string {
	return isPossible(&PossibleDoomConfigLocations)
}

// IsEmacs checks if emacs was found
func IsEmacs() *string {
	return isPossible(&PossibleEmacsConfigLocations)
}

func isPossible(s *[]string) *string {
	for _, loc := range *s {
		_, err := os.Stat(loc)
		if err == nil {
			return &loc
		}
	}
	return nil
}
