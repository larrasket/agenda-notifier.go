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

// IsInitialized checks if the package is intialized
func IsInitialized() bool {
	_, err := os.Stat(ConfigFileLoc)
	return err == nil
}
