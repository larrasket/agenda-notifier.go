package main

import (
	"fmt"
	. "github.com/salehmu/notifier.go/internal/notifiy"
	. "github.com/salehmu/notifier.go/pkg/reader"
)

func main() {
	config, err := InitializeReader()
	if err != nil {
		L.Fatal(fmt.Sprintf(`Couldn't initialize/read config: %s`, err))
	}
	ListenAndServe(config)
}
