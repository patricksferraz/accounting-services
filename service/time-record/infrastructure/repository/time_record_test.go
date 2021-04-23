package repository_test

import (
	"testing"
	"time"

	"github.com/patricksferraz/accounting-services/service/time-record/domain/model"
	"github.com/patricksferraz/accounting-services/service/time-record/infrastructure/db"
	"github.com/patricksferraz/accounting-services/service/time-record/infrastructure/repository"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func TestRepository_Register(t *testing.T) {

	_db, _ := db.ConnectMongoDB()
	repository := repository.NewTimeRecordRepositoryDb(_db)

	now := time.Now()
	description := faker.Lorem().Sentence(10)
	employeeID := uuid.NewV4().String()
	timeRecord, _ := model.NewTimeRecord(now, description, employeeID)

	err := repository.Register(timeRecord)
	require.Nil(t, err)
}

func TestRepository_Save(t *testing.T) {

	_db, _ := db.ConnectMongoDB()
	repository := repository.NewTimeRecordRepositoryDb(_db)

	now := time.Now()
	description := faker.Lorem().Sentence(10)
	employeeID := uuid.NewV4().String()
	timeRecord, _ := model.NewTimeRecord(now, description, employeeID)

	repository.Register(timeRecord)

	timeRecord.Description = faker.Lorem().Sentence(10)
	err := repository.Save(timeRecord)
	require.Nil(t, err)
}

func TestRepository_Find(t *testing.T) {

	_db, _ := db.ConnectMongoDB()
	repository := repository.NewTimeRecordRepositoryDb(_db)

	// NOTE: time.Time is in nanoseconds and mongodb in milliseconds
	y, m, d := time.Now().Date()
	now := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	description := faker.Lorem().Sentence(10)
	employeeID := uuid.NewV4().String()
	timeRecord, _ := model.NewTimeRecord(now, description, employeeID)

	repository.Register(timeRecord)

	timeRecordDb, err := repository.Find(timeRecord.ID)
	require.Nil(t, err)
	require.Equal(t, timeRecord.ID, timeRecordDb.ID)
	require.True(t, timeRecord.Time.Equal(timeRecordDb.Time))
	require.Equal(t, timeRecord.Status, timeRecordDb.Status)
	require.Equal(t, timeRecord.Description, timeRecordDb.Description)
	require.Equal(t, timeRecord.RegularTime, timeRecordDb.RegularTime)
	require.Equal(t, timeRecord.EmployeeID, timeRecordDb.EmployeeID)
	require.Equal(t, timeRecord.ApprovedBy, timeRecordDb.ApprovedBy)
	require.NotEmpty(t, timeRecordDb.CreatedAt)
	require.Empty(t, timeRecordDb.UpdatedAt)
}

func TestRepository_FindAllByEmployeeID(t *testing.T) {

	_db, _ := db.ConnectMongoDB()
	repository := repository.NewTimeRecordRepositoryDb(_db)

	// NOTE: time.Time is in nanoseconds and mongodb in milliseconds
	y, m, d := time.Now().Date()
	now := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	description := faker.Lorem().Sentence(10)
	employeeID := uuid.NewV4().String()
	timeRecord, _ := model.NewTimeRecord(now, description, employeeID)

	repository.Register(timeRecord)

	timeRecordsDb, err := repository.FindAllByEmployeeID(timeRecord.EmployeeID)
	require.Nil(t, err)
	require.Equal(t, timeRecord.ID, timeRecordsDb[0].ID)
	require.True(t, timeRecord.Time.Equal(timeRecordsDb[0].Time))
	require.Equal(t, timeRecord.Status, timeRecordsDb[0].Status)
	require.Equal(t, timeRecord.Description, timeRecordsDb[0].Description)
	require.Equal(t, timeRecord.RegularTime, timeRecordsDb[0].RegularTime)
	require.Equal(t, timeRecord.EmployeeID, timeRecordsDb[0].EmployeeID)
	require.Equal(t, timeRecord.ApprovedBy, timeRecordsDb[0].ApprovedBy)
	require.NotEmpty(t, timeRecordsDb[0].CreatedAt)
	require.Empty(t, timeRecordsDb[0].UpdatedAt)
}
