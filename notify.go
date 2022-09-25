package main

import (
	"time"
)

type Entity struct {
	Type, Name string
	Time       time.Time
}
