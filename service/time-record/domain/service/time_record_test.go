package service_test

import (
	"errors"
	"testing"
	"time"

	"github.com/patricksferraz/accounting-services/service/time-record/domain/model"
	"github.com/patricksferraz/accounting-services/service/time-record/domain/service"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

type repository struct{}

func (r *repository) Register(timeRecord *model.TimeRecord) error {
	// NOTE: Force error
	if timeRecord.Description == "error" {
		return errors.New("")
	}
	return nil
}

func (r *repository) Save(timeRecord *model.TimeRecord) error {
	// NOTE: Force error
	if timeRecord.ID == "c03c4cd4-5211-4209-ac68-17e441152b1d" {
		return errors.New("")
	}
	return nil
}

func (r *repository) Find(id string) (*model.TimeRecord, error) {
	timeRecord := model.TimeRecord{
		Time:        time.Now().AddDate(0, 0, -1),
		Status:      model.Pending,
		Description: faker.Lorem().Sentence(10),
		RegularTime: false,
		EmployeeID:  "67fe1eea-25a4-4f23-bf67-64f9a085311d",
	}
	timeRecord.ID = id
	// NOTE: Force error
	if id == "c4a80742-5294-4f1e-8ea9-5126c9389d6f" {
		return nil, errors.New("")
	}
	return &timeRecord, nil
}

func (r *repository) FindAllByEmployeeID(employeeID string) ([]*model.TimeRecord, error) {
	var timeRecords []*model.TimeRecord
	// NOTE: Force error
	if employeeID == "" {
		return nil, errors.New("")
	}
	return timeRecords, nil
}

func TestService_Register(t *testing.T) {

	timeRecordService := service.NewTimeRecordService(new(repository))

	_time := time.Now()
	description := faker.Lorem().Sentence(10)
	employeeID := uuid.NewV4().String()
	timeRecord, err := timeRecordService.Register(_time, description, employeeID)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(timeRecord.ID))
	require.Equal(t, timeRecord.Time, _time)
	require.Equal(t, timeRecord.Status, model.Approved)
	require.Equal(t, timeRecord.Description, description)
	require.Equal(t, timeRecord.RegularTime, true)
	require.Equal(t, timeRecord.EmployeeID, employeeID)

	_, err = timeRecordService.Register(_time.AddDate(0, 0, 1), description, employeeID)
	require.NotNil(t, err)
	_, err = timeRecordService.Register(_time, "error", employeeID)
	require.NotNil(t, err)
}

func TestService_Approve(t *testing.T) {

	timeRecordService := service.NewTimeRecordService(new(repository))

	id := uuid.NewV4().String()
	approvedBy := uuid.NewV4().String()
	timeRecord, err := timeRecordService.Approve(id, approvedBy)
	require.Nil(t, err)
	require.Equal(t, timeRecord.ApprovedBy, approvedBy)

	_, err = timeRecordService.Approve("c4a80742-5294-4f1e-8ea9-5126c9389d6f", approvedBy)
	require.NotNil(t, err)
	_, err = timeRecordService.Approve(id, "67fe1eea-25a4-4f23-bf67-64f9a085311d")
	require.NotNil(t, err)
	_, err = timeRecordService.Approve("c03c4cd4-5211-4209-ac68-17e441152b1d", approvedBy)
	require.NotNil(t, err)
}

func TestService_Find(t *testing.T) {

	timeRecordService := service.NewTimeRecordService(new(repository))

	id := uuid.NewV4().String()
	_, err := timeRecordService.Find(id)
	require.Nil(t, err)
	_, err = timeRecordService.Find("c4a80742-5294-4f1e-8ea9-5126c9389d6f")
	require.NotNil(t, err)
}

func TestService_FindAllByEmployeeID(t *testing.T) {

	timeRecordService := service.NewTimeRecordService(new(repository))

	employeeID := uuid.NewV4().String()
	_, err := timeRecordService.FindAllByEmployeeID(employeeID)
	require.Nil(t, err)
	_, err = timeRecordService.FindAllByEmployeeID("")
	require.NotNil(t, err)
}
