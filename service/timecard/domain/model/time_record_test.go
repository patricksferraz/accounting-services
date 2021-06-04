package model_test

import (
	"testing"
	"time"

	"github.com/patricksferraz/timecard-service/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func TestModel_NewTimeRecord(t *testing.T) {

	workedDayID := uuid.NewV4().String()
	createByID := uuid.NewV4().String()
	description := faker.Lorem().Sentence(10)

	_time := time.Now()
	timeRecord, err := model.NewTimeRecord(_time, description, workedDayID, createByID)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(timeRecord.ID))
	require.Equal(t, timeRecord.Time, _time)
	require.Equal(t, timeRecord.Description, description)
	require.Equal(t, timeRecord.WorkedDayID, workedDayID)
	require.Equal(t, timeRecord.CreatedBy, createByID)

	_, err = model.NewTimeRecord(time.Time{}, "", workedDayID, createByID)
	require.NotNil(t, err)
	_, err = model.NewTimeRecord(_time, "", "", createByID)
	require.NotNil(t, err)
	_, err = model.NewTimeRecord(_time, "", workedDayID, "")
	require.NotNil(t, err)
}
