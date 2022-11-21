package main

import (
	"bytes"
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
		if err != nil /* && errors.Is(err, &exec.ExitError{})  */ {
			fmt.Println(ReadDataErr)
			return
		}
		start := bytes.Index(data, []byte(AgendaStart))
		end := bytes.Index(data, []byte(AgendaEnd))
		data = data[start+len(AgendaStart) : end]
		coming, err := ComingEntity(data)
		if err != nil {
			fmt.Println(ReadDataErr)
			return
		}
		fmt.Println(coming.Name)
	}
}
