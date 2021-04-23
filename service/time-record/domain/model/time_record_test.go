package model_test

import (
	"testing"
	"time"

	"github.com/patricksferraz/accounting-services/service/time-record/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func TestModel_NewTimeRecord(t *testing.T) {

	now := time.Now()
	description := faker.Lorem().Sentence(10)
	employeeID := uuid.NewV4().String()
	timeRecord, err := model.NewTimeRecord(now, description, employeeID)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(timeRecord.ID))
	require.Equal(t, timeRecord.Time, now)
	require.Equal(t, timeRecord.Status, model.Approved)
	require.Equal(t, timeRecord.Description, description)
	require.Equal(t, timeRecord.RegularTime, true)
	require.Equal(t, timeRecord.EmployeeID, employeeID)

	yesterday := now.AddDate(0, 0, -1)
	timeRecord, err = model.NewTimeRecord(yesterday, description, employeeID)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(timeRecord.ID))
	require.Equal(t, timeRecord.Time, yesterday)
	require.Equal(t, timeRecord.Status, model.Pending)
	require.Equal(t, timeRecord.Description, description)
	require.Equal(t, timeRecord.RegularTime, false)
	require.Equal(t, timeRecord.EmployeeID, employeeID)

	tomorrow := now.AddDate(0, 0, 1)
	_, err = model.NewTimeRecord(tomorrow, description, employeeID)
	require.NotNil(t, err)
	_, err = model.NewTimeRecord(time.Time{}, "", employeeID)
	require.NotNil(t, err)
	_, err = model.NewTimeRecord(tomorrow, "", employeeID)
	require.NotNil(t, err)
	_, err = model.NewTimeRecord(tomorrow, description, "")
	require.NotNil(t, err)
}

func TestModel_ChangeStatusOfATimeRecord(t *testing.T) {

	yesterday := time.Now().AddDate(0, 0, -1)
	description := faker.Lorem().Sentence(10)
	employeeID := uuid.NewV4().String()
	timeRecord, _ := model.NewTimeRecord(yesterday, description, employeeID)

	approvedBy := uuid.NewV4().String()
	err := timeRecord.Approve(approvedBy)
	require.Nil(t, err)
	require.Equal(t, timeRecord.Status, model.Approved)
	require.True(t, timeRecord.CreatedAt.Before(timeRecord.UpdatedAt))

	timeRecord, _ = model.NewTimeRecord(yesterday, description, employeeID)

	err = timeRecord.Approve(employeeID)
	require.NotNil(t, err)
}
