package repository

import (
	"github.com/patricksferraz/timecard-service/domain/model"
)

type TimecardRepositoryInterface interface {
	Register(timecard *model.Timecard) error
	Save(timecard *model.Timecard) error
	Find(id string) (*model.Timecard, error)
	FindAllByCompanyID(companyID string) ([]*model.Timecard, error)
	FindAllByEmployeeID(employeeID string) ([]*model.Timecard, error)
}
