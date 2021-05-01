package repository

import (
	"context"
	"time"

	"github.com/patricksferraz/accounting-services/service/time-record/domain/model"
)

type TimeRecordRepositoryInterface interface {
	Register(ctx context.Context, timeRecord *model.TimeRecord) error
	Save(ctx context.Context, timeRecord *model.TimeRecord) error
	Find(ctx context.Context, id string) (*model.TimeRecord, error)
	FindAllByEmployeeID(ctx context.Context, employeeID string, fromDate, toDate time.Time) ([]*model.TimeRecord, error)
}
