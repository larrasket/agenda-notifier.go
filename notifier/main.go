package main

import (
	"fmt"
	"github.com/salehmu/notifier.go/internal/notifiy"
	. "github.com/salehmu/notifier.go/pkg/reader"
)

func main() {
	logger := NewLogger()
	config, err := InitializeReader()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Couldn't read configuration file %s", err))
	}
	notifiy.ListenAndServe(config)
}
