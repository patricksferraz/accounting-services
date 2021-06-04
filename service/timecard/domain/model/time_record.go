package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type TimeRecord struct {
	Base        `bson:",inline" valid:"-"`
	Time        time.Time `bson:"time" valid:"required"`
	Description string    `bson:"description" valid:"-"`
	WorkedDayID string    `bson:"worked_day_id" valid:"uuid"`
	CreatedBy   string    `bson:"created_by" valid:"uuid"`
}

func (p *TimeRecord) isValid() error {
	_, err := govalidator.ValidateStruct(p)
	return err
}

func NewTimeRecord(_time time.Time, description string, workedDayID string, createdBy string) (*TimeRecord, error) {

	timeRecord := TimeRecord{
		Time:        _time,
		Description: description,
		WorkedDayID: workedDayID,
		CreatedBy:   createdBy,
	}

	timeRecord.ID = uuid.NewV4().String()
	timeRecord.CreatedAt = time.Now()

	err := timeRecord.isValid()
	if err != nil {
		return nil, err
	}

	return &timeRecord, nil
}
