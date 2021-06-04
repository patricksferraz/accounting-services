package service

import (
	"time"

	"github.com/patricksferraz/timecard-service/domain/model"
	"github.com/patricksferraz/timecard-service/domain/repository"
)

type TimecardService struct {
	TimecardRepository   repository.TimecardRepositoryInterface
	TimeRecordRepository repository.TimeRecordRepositoryInterface
}

func (p *TimecardService) RegisterTimecard(date time.Time, companyID string, employeeID string) (*model.Timecard, error) {
	timecard, err := model.NewTimecard(date, companyID, employeeID)
	if err != nil {
		return nil, err
	}

	err = p.TimecardRepository.Register(timecard)
	if err != nil {
		return nil, err
	}

	return timecard, nil
}

func (p *TimecardService) RegisterTimeRecord(_time time.Time, timecardID string) (*model.TimeRecord, error) {
	timecard, err := p.TimecardRepository.Find(timecardID)
	if err != nil {
		return nil, err
	}

	timeRecord, err := model.NewTimeRecord(_time, timecard.ID)
	if err != nil {
		return nil, err
	}

	// TODO: Adds transaction
	timecard.AddTimeRecord(timeRecord.ID)
	err = p.TimecardRepository.Save(timecard)
	if err != nil {
		return nil, err
	}

	err = p.TimeRecordRepository.Register(timeRecord)
	if err != nil {
		return nil, err
	}

	return timeRecord, nil
}

func (p *TimecardService) WaitConfirmationTimecard(timecardID string) (*model.Timecard, error) {
	timecard, err := p.TimecardRepository.Find(timecardID)
	if err != nil {
		return nil, err
	}

	timecard.WaitConfirmation()
	err = p.TimecardRepository.Save(timecard)
	if err != nil {
		return nil, err
	}

	return timecard, nil
}

func (p *TimecardService) ConfirmTimecard(timecardID string) (*model.Timecard, error) {
	timecard, err := p.TimecardRepository.Find(timecardID)
	if err != nil {
		return nil, err
	}

	timecard.Confirm()
	err = p.TimecardRepository.Save(timecard)
	if err != nil {
		return nil, err
	}

	return timecard, nil
}

func (p *TimecardService) FindTimecard(timecardID string) (*model.Timecard, error) {
	timecard, err := p.TimecardRepository.Find(timecardID)
	if err != nil {
		return nil, err
	}
	return timecard, nil
}

func (p *TimecardService) FindAllTimecardByCompanyID(companyID string) ([]*model.Timecard, error) {
	timecards, err := p.TimecardRepository.FindAllByCompanyID(companyID)
	if err != nil {
		return nil, err
	}
	return timecards, nil
}

func (p *TimecardService) FindAllTimecardByEmployeeID(employeeID string) ([]*model.Timecard, error) {
	timecards, err := p.TimecardRepository.FindAllByEmployeeID(employeeID)
	if err != nil {
		return nil, err
	}
	return timecards, nil
}

func (p *TimecardService) FindAllTimeRecordByTimecardID(timecardID string) ([]*model.TimeRecord, error) {
	timeRecords, err := p.TimeRecordRepository.FindAllByTimecardID(timecardID)
	if err != nil {
		return nil, err
	}
	return timeRecords, nil
}
