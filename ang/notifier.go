package main

import (
	"fmt"
	. "github.com/larrasket/notifier.go/internal/notifiy"
	. "github.com/larrasket/notifier.go/pkg/reader"
)

func main() {
	config, err := InitializeReader()
	if err != nil {
		L.Fatal(fmt.Sprintf(`Couldn't initialize/read config: %s`, err))
	}
	ListenAndServe(config)
}
