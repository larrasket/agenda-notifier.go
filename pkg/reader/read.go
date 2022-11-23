package reader

import (
	"bytes"
	"encoding/csv"
	"strings"
	"time"
)

type Entity struct {
	Type, Name, Label string
	Time              time.Time
}

func InitializeReader() (*Config, error) {
	if !IsInitialized() {
		_, err := InitConfig()
		if err != nil {
			return nil, err
		}
	}
	config, err := ReadConfig()
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func ComingEntity(data []byte) (*Entity, error) {
	entities, err := readData(data)
	if err != nil {
		return nil, err
	}
	t := Entity{}
	nows := time.Now().Format(TimeFormat)
	now, _ := time.Parse(TimeFormat, nows)
	for _, entity := range entities {
		if len(entity[0]) == 0 || len(entity[6]) == 0 {
			continue
		}
		clock, err := extractTime(entity)
		t.Type = entity[3]

		if err == nil && now.Before(clock) {
			t.Label = entity[0]
			t.Name = entity[1]
			t.Time = clock
			return &t, err
		} else if err != nil {
			return nil, err
		}
	}
	return nil, NoEntityErr
}

// TODO Skip done/killed entities
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
	t, err := time.Parse(TimeFormat, fullDate)
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
