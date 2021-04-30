package service

import (
	"time"

	"github.com/patricksferraz/accounting-services/service/time-record/domain/model"
	"github.com/patricksferraz/accounting-services/service/time-record/domain/repository"
)

type TimeRecordService struct {
	TimeRecordRepository repository.TimeRecordRepositoryInterface
}

func (p *TimeRecordService) Register(_time time.Time, description, employeeID string) (*model.TimeRecord, error) {
	timeRecord, err := model.NewTimeRecord(_time, description, employeeID)
	if err != nil {
		return nil, err
	}

	err = p.TimeRecordRepository.Register(timeRecord)
	if err != nil {
		return nil, err
	}

	return timeRecord, nil
}

func (p *TimeRecordService) Approve(id, employeeID string) (*model.TimeRecord, error) {
	timeRecord, err := p.TimeRecordRepository.Find(id)
	if err != nil {
		return nil, err
	}

	err = timeRecord.Approve(employeeID)
	if err != nil {
		return nil, err
	}

	err = p.TimeRecordRepository.Save(timeRecord)
	if err != nil {
		return nil, err
	}

	return timeRecord, nil
}

func (p *TimeRecordService) Find(id string) (*model.TimeRecord, error) {
	timeRecord, err := p.TimeRecordRepository.Find(id)
	if err != nil {
		return nil, err
	}
	return timeRecord, nil
}

func (p *TimeRecordService) FindAllByEmployeeID(employeeID string, fromDate, toDate time.Time) ([]*model.TimeRecord, error) {
	timeRecords, err := p.TimeRecordRepository.FindAllByEmployeeID(employeeID, fromDate, toDate)
	if err != nil {
		return nil, err
	}
	return timeRecords, nil
}

func NewTimeRecordService(timeRecordRepository repository.TimeRecordRepositoryInterface) *TimeRecordService {
	return &TimeRecordService{
		TimeRecordRepository: timeRecordRepository,
	}
}
