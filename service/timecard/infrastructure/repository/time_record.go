package repository

import (
	"github.com/patricksferraz/timecard-service/domain/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	timeRecordCollection string = "time_record"
)

type TimeRecordRepositoryDb struct {
	Db *mgo.Database
}

func (p *TimeRecordRepositoryDb) Register(timeRecord *model.TimeRecord) error {
	err := p.Db.C(timeRecordCollection).Insert(timeRecord)
	return err
}

func (p *TimeRecordRepositoryDb) Save(timeRecord *model.TimeRecord) error {
	err := p.Db.C(timeRecordCollection).UpdateId(timeRecord.ID, timeRecord)
	return err
}

func (p *TimeRecordRepositoryDb) Find(id string) (*model.TimeRecord, error) {
	var timeRecord *model.TimeRecord
	err := p.Db.C(timeRecordCollection).FindId(id).One(&timeRecord)
	return timeRecord, err
}

func (p *TimeRecordRepositoryDb) FindAllByTimecardID(timecardID string) ([]*model.TimeRecord, error) {
	var timeRecords []*model.TimeRecord
	err := p.Db.C(timeRecordCollection).Find(bson.M{"timecard_id": timecardID}).All(&timeRecords)
	return timeRecords, err
}
