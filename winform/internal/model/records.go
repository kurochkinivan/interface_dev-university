package model

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Record struct {
	PlayerName string
	Score      int
	TimeSpent  time.Duration
	DatePlayed time.Time
	Level      int
}

type StatisticsSubscriber struct {
	records []Record
	callback func([]Record)
}

func (s *StatisticsSubscriber) Update() {
	settings := GetSettings()
	newRecords, err := ReadRecords(settings.PathToStatistics)
	if err != nil {
		fmt.Printf("Ошибка при обновлении статистики: %v\n", err)
		return
	}
	s.records = newRecords
	fmt.Println("updating, new records:", newRecords)
	
	callback := s.callback
	if callback != nil {
		callback(newRecords)
	}
}

func (s *StatisticsSubscriber) SetCallBack(cb func ([]Record)) {
	s.callback = cb
}

func (s *StatisticsSubscriber) GetRecords() []Record {
	fmt.Println(len(s.records))
	return s.records
}

func ReadRecords(pathToFile string) ([]Record, error) {
	const op string = "model.ReadRecords"

	f, err := os.Open(makePathToData(pathToFile))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer f.Close()

	var records []Record
	if err := json.NewDecoder(f).Decode(&records); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return records, nil
}

func makePathToData(pathToFile string) string {
	wd, _ := os.Getwd()
	wd = filepath.Dir(filepath.Dir(wd))
	return filepath.Join(wd, "data", pathToFile)
}
