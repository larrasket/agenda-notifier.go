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
	now, _ := time.Parse(timeFormat, time.Now().Format(timeFormat))
	for _, entity := range entities {
		if len(entity[0]) == 0 {
			continue
		}
		entitytime, xrr := extractTime(entity)
		if xrr == nil && entitytime.After(now) {
			t.Type = entity[0]
			t.Name = entity[1]
			t.Time = entitytime
			return
		}
	}
	err = ErrNoUpcomming
	return
}

func extractTime(entity []string) (time.Time, error) {
	stime := entity[6]
	if strings.ContainsRune(stime, '-') {
		splits := strings.Split(stime, "-")
		stime = splits[0]
	} else if strings.ContainsRune(stime, '.') {
		splits := strings.Split(stime, ".")
		stime = splits[0]
	}
	t, err := time.Parse(timeFormat, stime)
	return t, err
}

func readData(data []byte) ([][]string, error) {

	r := csv.NewReader(bytes.NewReader(data))

	// skip first line
	//if _, err := r.Read(); err != nil {
	//	return [][]string{}, err
	//}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}
	return records, nil
}
