package controller

import (
	"github.com/kurochkinivan/winform/internal/model"
)

func NewStatisticsSubscriber() *model.StatisticsSubscriber { 
	statisticsSubscriber := &model.StatisticsSubscriber{}
	model.GetSettings().Attach(statisticsSubscriber)
	
	statisticsSubscriber.Update()

	return statisticsSubscriber
}

func GetRecords(pathToFile string) ([]model.Record, error) {
	return model.ReadRecords(pathToFile)
}
