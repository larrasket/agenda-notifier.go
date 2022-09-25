package main

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

func main() {
	if !IsInitialized() {
		_, err := InitConfig()
		if err != nil {
			fmt.Println(WriteConfErr)
			return
		}
	}
	config, err := ReadConfig()
	if err != nil {
		fmt.Println(ReadConfErr)
		return
	}

	if config.Doom {
		data, err := exec.Command(config.DoomScript, ExportScriptLoc).Output()
		if err != nil && errors.Is(err, &exec.ExitError{}) {
			fmt.Println(ReadDataErr)
			return
		}

		start := bytes.Index(data, []byte(AgendaStart))
		end := bytes.Index(data, []byte(AgendaEnd))
		//fmt.Println(string())

		coming, err := ComingEntity(data[start+len(AgendaStart) : end])
		if err != nil {
			fmt.Println(ReadDataErr)
			return
		}

		fmt.Println(coming.Name)
	} else {

	}

}
