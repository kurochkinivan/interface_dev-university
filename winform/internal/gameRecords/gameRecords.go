package gamerecords

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type GameRecord struct {
	PlayerName string
	Score      int
	TimeSpent  time.Duration
	DatePlayed time.Time
	Level      int
}

func ReadDataFile(pathToFile string) ([]GameRecord, error) {
	const op string = "records.readDataFile"

	f, err := os.Open(makePathToData(pathToFile))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var records []GameRecord
	err = json.Unmarshal(data, &records)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return records, nil
}

func makePathToData(pathToFile string) string {
	wd, _ := os.Getwd()
	wd = filepath.Dir(filepath.Dir(wd))
	return filepath.Join(wd, "data", pathToFile)
}
