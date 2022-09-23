package main

import "os"

// IsDoom checks if doom emacs is installed
func IsDoom() (bool, *string) {
	return isPossible(&PossibleDoomConfigLocations)
}

// IsEmacs checks if emacs is found
func IsEmacs() (bool, *string) {
	return isPossible(&PossibleEmacsConfigLocations)
}

func isPossible(s *[]string) (bool, *string) {
	for _, loc := range *s {
		_, err := os.Stat(loc)
		if err != nil {
			return true, &loc
		}
	}
	return false, nil
}

// IsIntialized checks if the package is intialized
func IsIntialized() bool {
	_, err := os.Stat(ConfigDirDef)
	return err != nil
}
