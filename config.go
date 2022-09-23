package main

// Type to generate a configuration file.
type config struct {
	Doom               bool   `json:"doom"`                // whether user uses doom or not.
	EmacsDir           string `json:"emacsDir"`            // emacs directory 	   		           (if !doom)
	ScanInt            int    `json:scanInt`               // scan interval.                        (in second)
	BeforeNotification int    `json:beforeNotificationdef` // time before notifing                  (in second)
	NotifyCommand      string `json:notifiyCommand`        // command to use for sending notification
}

// GetConfig returns default configuration data
func MakeConfig(doom bool, dir string) config {
	conf := config{
		doom,
		dir,
		ScanIntDef,
		BeforeNotificationDef,
		NotifyCommandDef,
	}
	return conf
}
