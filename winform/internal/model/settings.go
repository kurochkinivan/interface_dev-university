package model

import (
	"sync"
)

type Settings struct {
	PathToStatistics string
	subcribers       []Subscriber
}

var (
	SettingsInstance *Settings
	once             sync.Once
)

func GetSettings() *Settings {
	once.Do(func() {
		SettingsInstance = &Settings{
			PathToStatistics: "default.json",
		}
	})
	return SettingsInstance
}

func (s *Settings) UpdatePath(newPath string) {
	s.PathToStatistics = newPath
	s.Notify()
}

func (s *Settings) Attach(sub Subscriber) {
	s.subcribers = append(s.subcribers, sub)
}

func (s *Settings) Detach(subToDel Subscriber) {
	for i, sub := range s.subcribers {
		if sub == subToDel {
			s.subcribers = append(s.subcribers[:i], s.subcribers[i+1:]...)
			break
		}
	}
}

func (s *Settings) Notify() {
	for _, sub := range s.subcribers {
		sub.Update()
	}
}
