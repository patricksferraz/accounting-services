package repository

import (
	"github.com/patricksferraz/timecard-service/domain/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	timecardCollection string = "timecard"
)

type TimecardRepositoryDb struct {
	Db *mgo.Database
}

func (p *TimecardRepositoryDb) Register(timecard *model.Timecard) error {
	err := p.Db.C(timecardCollection).Insert(timecard)
	return err
}

func (p *TimecardRepositoryDb) Save(timecard *model.Timecard) error {
	err := p.Db.C(timecardCollection).UpdateId(timecard.ID, timecard)
	return err
}

func (p *TimecardRepositoryDb) Find(id string) (*model.Timecard, error) {
	var timecard *model.Timecard
	err := p.Db.C(timecardCollection).FindId(id).One(&timecard)
	return timecard, err
}

func (p *TimecardRepositoryDb) FindAllByCompanyID(companyID string) ([]*model.Timecard, error) {
	var timecards []*model.Timecard
	err := p.Db.C(timecardCollection).Find(bson.M{"company_id": companyID}).All(&timecards)
	return timecards, err
}

func (p *TimecardRepositoryDb) FindAllByEmployeeID(employeeID string) ([]*model.Timecard, error) {
	var timecards []*model.Timecard
	err := p.Db.C(timecardCollection).Find(bson.M{"employee_id": employeeID}).All(&timecards)
	return timecards, err
}
