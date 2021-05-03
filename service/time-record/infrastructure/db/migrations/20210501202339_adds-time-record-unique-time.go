package migrations

import (
	"github.com/patricksferraz/accounting-services/service/time-record/infrastructure/repository"
	migrate "github.com/patricksferraz/mongo-migrate"
	"gopkg.in/mgo.v2"
)

func init() {
	index := mgo.Index{
		Name:     "time_record_id_time_unique",
		Key:      []string{"employee_id", "time"},
		Unique:   true,
		DropDups: true,
	}
	migrate.Register(func(db *mgo.Database) error {
		return db.C(repository.TimeRecordCollection).EnsureIndex(index)
	}, func(db *mgo.Database) error {
		return db.C(repository.TimeRecordCollection).DropIndexName(index.Name)
	})
}
