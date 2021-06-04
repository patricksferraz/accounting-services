package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type WorkedDay struct {
	Base        `bson:",inline" valid:"-"`
	Date        time.Time     `bson:"date" valid:"-"`
	TimeRecords []*TimeRecord `bson:"time_records" valid:"-"`
	TimecardID  string        `bson:"timecard_id" valid:"uuid"`
}

func (w *WorkedDay) isValid() error {
	_, err := govalidator.ValidateStruct(w)
	return err
}

func NewWorkedDay(date time.Time, companyID string, employeeID string) (*Timecard, error) {

	timecard := Timecard{
		Status:     Open,
		Date:       date,
		CompanyID:  companyID,
		EmployeeID: employeeID,
	}

	timecard.ID = uuid.NewV4().String()
	timecard.CreatedAt = time.Now()

	err := timecard.isValid()
	if err != nil {
		return nil, err
	}

	return &timecard, nil
}
