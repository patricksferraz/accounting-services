package service

import (
	"context"
	"time"

	"github.com/c4ut/accounting-services/service/time-record/domain/model"
	"github.com/c4ut/accounting-services/service/time-record/domain/repository"
	"go.elastic.co/apm"
)

type TimeRecordService struct {
	TimeRecordRepository repository.TimeRecordRepositoryInterface
}

func (p *TimeRecordService) Register(ctx context.Context, _time time.Time, description, employeeID string) (*model.TimeRecord, error) {
	span, ctx := apm.StartSpan(ctx, "Register", "service")
	defer span.End()

	timeRecord, err := model.NewTimeRecord(_time, description, employeeID)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		return nil, err
	}

	err = p.TimeRecordRepository.Register(ctx, timeRecord)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		return nil, err
	}

	return timeRecord, nil
}

func (p *TimeRecordService) Approve(ctx context.Context, id, employeeID string) (*model.TimeRecord, error) {
	span, ctx := apm.StartSpan(ctx, "Approve", "service")
	defer span.End()

	timeRecord, err := p.TimeRecordRepository.Find(ctx, id)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		return nil, err
	}

	err = timeRecord.Approve(employeeID)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		return nil, err
	}

	err = p.TimeRecordRepository.Save(ctx, timeRecord)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		return nil, err
	}

	return timeRecord, nil
}

func (p *TimeRecordService) Refuse(ctx context.Context, id, refusedReason, employeeID string) (*model.TimeRecord, error) {
	span, ctx := apm.StartSpan(ctx, "Refuse", "service")
	defer span.End()

	timeRecord, err := p.TimeRecordRepository.Find(ctx, id)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		return nil, err
	}

	err = timeRecord.Refuse(employeeID, refusedReason)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		return nil, err
	}

	err = p.TimeRecordRepository.Save(ctx, timeRecord)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		return nil, err
	}

	return timeRecord, nil
}

func (p *TimeRecordService) Find(ctx context.Context, id string) (*model.TimeRecord, error) {
	span, ctx := apm.StartSpan(ctx, "Find", "service")
	defer span.End()

	timeRecord, err := p.TimeRecordRepository.Find(ctx, id)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		return nil, err
	}
	return timeRecord, nil
}

func (p *TimeRecordService) FindAllByEmployeeID(ctx context.Context, employeeID string, fromDate, toDate time.Time) ([]*model.TimeRecord, error) {
	span, ctx := apm.StartSpan(ctx, "FindAllByEmployeeID", "service")
	defer span.End()

	timeRecords, err := p.TimeRecordRepository.FindAllByEmployeeID(ctx, employeeID, fromDate, toDate)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		return nil, err
	}
	return timeRecords, nil
}

func NewTimeRecordService(timeRecordRepository repository.TimeRecordRepositoryInterface) *TimeRecordService {
	return &TimeRecordService{
		TimeRecordRepository: timeRecordRepository,
	}
}
