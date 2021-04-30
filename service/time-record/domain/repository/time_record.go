package repository

import (
	"time"

	"github.com/patricksferraz/accounting-services/service/time-record/domain/model"
)

type TimeRecordRepositoryInterface interface {
	Register(timeRecord *model.TimeRecord) error
	Save(timeRecord *model.TimeRecord) error
	Find(id string) (*model.TimeRecord, error)
	FindAllByEmployeeID(employeeID string, fromDate, toDate time.Time) ([]*model.TimeRecord, error)
}
