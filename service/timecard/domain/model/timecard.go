package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Timecard struct {
	Base          `bson:",inline" valid:"-"`
	Status        TimecardStatus `bson:"status" valid:"timecardStatus"`
	StartDate     time.Time      `bson:"start_date" valid:"required"`
	EndDate       time.Time      `bson:"end_date" valid:"required"`
	WorkedHours   time.Duration  `bson:"worked_hours" valid:"-"`
	OvertimeHours time.Duration  `bson:"overtime_hours" valid:"-"`
	AbsenceHours  time.Duration  `bson:"absence_hours" valid:"-"`
	WorkedDays    []*WorkedDay   `bson:"work_days" valid:"-"`
	CompanyID     string         `bson:"company_id" valid:"uuid"`
	EmployeeID    string         `bson:"employee_id" valid:"uuid"`
	AuditedBy     string         `bson:"audited_by" valid:"-"`
	ConfirmedBy   string         `bson:"confirmed_by" valid:"-"`
}

func (t *Timecard) isValid() error {
	_, err := govalidator.ValidateStruct(t)
	return err
}

func (t *Timecard) WaitConfirmation(employeeID string) error {
	t.Status = WaitingConfirmation
	t.AuditedBy = employeeID
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Timecard) Confirm(employeeID string) error {
	t.Status = Confirmed
	t.ConfirmedBy = employeeID
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Timecard) AddWorkedDay(workedDay *WorkedDay) error {

	if t.Status == Confirmed {
		return errors.New("the timecard should not be changed after confirmed")
	}

	t.WorkedDays = append(t.WorkedDays, workedDay)
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func NewTimecard(startDate time.Time, endDate time.Time, workedDays []*WorkedDay, companyID string, employeeID string) (*Timecard, error) {

	timecard := Timecard{
		Status:     Open,
		StartDate:  startDate,
		EndDate:    endDate,
		WorkedDays: workedDays,
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
