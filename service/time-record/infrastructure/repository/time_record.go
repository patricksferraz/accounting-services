package repository

import (
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

func (t *TimeRecordRepositoryDb) Register(timeRecord *model.TimeRecord) error {
	err := t.Db.C(TimeRecordCollection).Insert(timeRecord)
	return err
}

func (t *TimeRecordRepositoryDb) Save(timeRecord *model.TimeRecord) error {
	err := t.Db.C(TimeRecordCollection).UpdateId(timeRecord.ID, timeRecord)
	return err
}

func (t *TimeRecordRepositoryDb) Find(id string) (*model.TimeRecord, error) {
	var timeRecord *model.TimeRecord
	err := t.Db.C(TimeRecordCollection).FindId(id).One(&timeRecord)
	return timeRecord, err
}

func (t *TimeRecordRepositoryDb) FindAllByEmployeeID(employeeID string) ([]*model.TimeRecord, error) {
	var timeRecords []*model.TimeRecord
	err := t.Db.C(TimeRecordCollection).Find(bson.M{"employee_id": employeeID}).Sort("-time").All(&timeRecords)
	return timeRecords, err
}

func NewTimeRecordRepositoryDb(database *mgo.Database) *TimeRecordRepositoryDb {
	return &TimeRecordRepositoryDb{
		Db: database,
	}
}
