package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/patricksferraz/accounting-services/utils"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type TimeRecord struct {
	Base        `bson:",inline" valid:"-"`
	Time        time.Time        `json:"time,omitempty" bson:"time" valid:"required"`
	Status      TimeRecordStatus `json:"status,omitempty" bson:"status" valid:"timeRecordStatus"`
	Description string           `json:"description,omitempty" bson:"description,omitempty" valid:"-"`
	RegularTime bool             `json:"regular_time,omitempty" bson:"regular_time" valid:"-"`
	EmployeeID  string           `json:"employee_id,omitempty" bson:"employee_id" valid:"uuid"`
	ApprovedBy  string           `json:"approved_by,omitempty" bson:"approved_by,omitempty" valid:"-"`
}

func (t *TimeRecord) isValid() error {

	if t.Time.After(time.Now()) {
		return errors.New("the registration time must not be longer than the current time")
	}

	if t.EmployeeID == t.ApprovedBy {
		return errors.New("the employee who recorded the time cannot be the same person who approves")
	}

	if !t.RegularTime && t.Description == "" {
		return errors.New("the description must not be empty when the registration is done in an irregular period")
	}

	_, err := govalidator.ValidateStruct(t)
	return err
}

func (t *TimeRecord) Approve(approvedBy string) error {

	if t.Status == Approved {
		return errors.New("the time record cannot be approved again")
	}

	t.Status = Approved
	t.UpdatedAt = time.Now()
	t.ApprovedBy = approvedBy
	err := t.isValid()
	return err
}

func NewTimeRecord(_time time.Time, description, employeeID string) (*TimeRecord, error) {

	timeRecord := TimeRecord{
		Time:        _time,
		Status:      Approved,
		Description: description,
		RegularTime: true,
		EmployeeID:  employeeID,
	}

	if !utils.DateEqual(_time, time.Now()) {
		timeRecord.Status = Pending
		timeRecord.RegularTime = false
	}

	timeRecord.ID = uuid.NewV4().String()
	timeRecord.CreatedAt = time.Now()

	err := timeRecord.isValid()
	if err != nil {
		return nil, err
	}

	return &timeRecord, nil
}
