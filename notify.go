package main

import (
	"time"
)

type Entity struct {
	Type, Name, Label string
	Time              time.Time
}
