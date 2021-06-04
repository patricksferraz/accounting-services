package model_test

import (
	"testing"
	"time"

	"github.com/patricksferraz/timecard-service/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestModel_NewTimecard(t *testing.T) {

	companyID := uuid.NewV4().String()
	employeeID := uuid.NewV4().String()

	date := time.Now()
	timecard, err := model.NewTimecard(date, companyID, employeeID)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(timecard.ID))
	require.Equal(t, timecard.Status, model.Open)
	require.Equal(t, timecard.Date, date)
	require.Equal(t, timecard.CompanyID, companyID)
	require.Equal(t, timecard.EmployeeID, employeeID)

	_, err = model.NewTimecard(time.Time{}, companyID, employeeID)
	require.NotNil(t, err)
	_, err = model.NewTimecard(date, "", employeeID)
	require.NotNil(t, err)
	_, err = model.NewTimecard(date, companyID, "")
	require.NotNil(t, err)
}

func TestModel_ChangeStatusOfATimecard(t *testing.T) {

	companyID := uuid.NewV4().String()
	employeeID := uuid.NewV4().String()

	date := time.Now()
	timecard, _ := model.NewTimecard(date, companyID, employeeID)

	err := timecard.WaitConfirmation()
	require.Nil(t, err)
	require.Equal(t, timecard.Status, model.WaitingConfirmation)
	require.True(t, timecard.CreatedAt.Before(timecard.UpdatedAt))

	err = timecard.Confirm()
	require.Nil(t, err)
	require.Equal(t, timecard.Status, model.Confirmed)
	require.True(t, timecard.CreatedAt.Before(timecard.UpdatedAt))
}

func TestModel_AddTimeRecordInATimecard(t *testing.T) {

	companyID := uuid.NewV4().String()
	employeeID := uuid.NewV4().String()

	date := time.Now()
	timecard, _ := model.NewTimecard(date, companyID, employeeID)

	timeRecordID := uuid.NewV4().String()
	err := timecard.AddTimeRecord(timeRecordID)
	require.Nil(t, err)
	require.Equal(t, timecard.Status, model.Open)
	require.Contains(t, timecard.TimeRecords, timeRecordID)
	require.True(t, timecard.CreatedAt.Before(timecard.UpdatedAt))

	timecard.Confirm()
	timeRecordID = uuid.NewV4().String()
	err = timecard.AddTimeRecord(timeRecordID)
	require.NotNil(t, err)
}
