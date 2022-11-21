package main

import (
	"bytes"
	"encoding/csv"
	"strings"
	"time"
)

func ComingEntity(data []byte) (t Entity, err error) {
	entities, err := readData(data)
	if err != nil {
		return
	}
	now := time.Now()
	for _, entity := range entities {
		if len(entity[0]) == 0 {
			continue
		}
		clock, xrr := extractTime(entity)
		t.Type = entity[3]
		if xrr == nil && clock.After(now) {
			t.Label = entity[0]
			t.Name = entity[1]
			t.Time = clock
			return
		}
		// TODO Handle Err
	}
	err = ErrNoUpcomming
	return
}

func extractTime(entity []string) (time.Time, error) {
	hours, days := entity[6], entity[5]
	if strings.ContainsRune(hours, '-') {
		splits := strings.Split(hours, "-")
		hours = splits[0]
	} else if strings.ContainsRune(hours, '.') {
		splits := strings.Split(hours, ".")
		hours = splits[0]
	}
	fullDate := days + "_" + hours
	t, err := time.Parse(timeFormat, fullDate)
	return t, err
}

func readData(data []byte) ([][]string, error) {
	r := csv.NewReader(bytes.NewReader(data))
	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return records, nil
}
