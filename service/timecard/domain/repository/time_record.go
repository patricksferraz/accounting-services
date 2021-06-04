package repository

import (
	"github.com/patricksferraz/timecard-service/domain/model"
)

type TimeRecordRepositoryInterface interface {
	Register(timeRecord *model.TimeRecord) error
	Save(timeRecord *model.TimeRecord) error
	Find(id string) (*model.TimeRecord, error)
	FindAllByTimecardID(timecardID string) ([]*model.TimeRecord, error)
}
