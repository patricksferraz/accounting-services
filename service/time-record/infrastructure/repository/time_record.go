package repository

import (
	"context"
	"time"

	"github.com/patricksferraz/accounting-services/service/time-record/domain/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	TimeRecordCollection string = "time_record"
)

type TimeRecordRepositoryDb struct {
	Db *mgo.Database
}

func (t *TimeRecordRepositoryDb) Register(ctx context.Context, timeRecord *model.TimeRecord) error {
	err := t.Db.C(TimeRecordCollection).Insert(timeRecord)
	return err
}

func (t *TimeRecordRepositoryDb) Save(ctx context.Context, timeRecord *model.TimeRecord) error {
	err := t.Db.C(TimeRecordCollection).UpdateId(timeRecord.ID, timeRecord)
	return err
}

func (t *TimeRecordRepositoryDb) Find(ctx context.Context, id string) (*model.TimeRecord, error) {
	var timeRecord *model.TimeRecord
	err := t.Db.C(TimeRecordCollection).FindId(id).One(&timeRecord)
	return timeRecord, err
}

func (t *TimeRecordRepositoryDb) FindAllByEmployeeID(ctx context.Context, employeeID string, fromDate, toDate time.Time) ([]*model.TimeRecord, error) {
	var timeRecords []*model.TimeRecord
	err := t.Db.C(TimeRecordCollection).Find(
		bson.M{
			"employee_id": employeeID,
			"time": bson.M{
				"$gte": fromDate,
				"$lte": toDate,
			},
		},
	).Sort("-time").All(&timeRecords)
	return timeRecords, err
}

func NewTimeRecordRepositoryDb(database *mgo.Database) *TimeRecordRepositoryDb {
	return &TimeRecordRepositoryDb{
		Db: database,
	}
}
