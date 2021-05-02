package migrations

import (
	"github.com/patricksferraz/accounting-services/service/time-record/infrastructure/repository"
	migrate "github.com/patricksferraz/mongo-migrate"
	"gopkg.in/mgo.v2"
)

func init() {
	migrate.Register(func(db *mgo.Database) error {
		return db.C(repository.TimeRecordCollection).EnsureIndex(
			mgo.Index{
				Name:   "time-record-unique-time",
				Key:    []string{"employee_id", "time"},
				Unique: true,
			},
		)
	}, func(db *mgo.Database) error {
		return db.C(repository.TimeRecordCollection).DropIndexName("time-record-unique-time")
	})
}
