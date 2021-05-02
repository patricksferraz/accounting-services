package db

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	_ "github.com/patricksferraz/accounting-services/service/time-record/infrastructure/db/migrations"
	"github.com/patricksferraz/accounting-services/utils"
	migrate "github.com/patricksferraz/mongo-migrate"
	mgo "gopkg.in/mgo.v2"
)

// TODO: Adds cover test
func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	if os.Getenv("ENV") == "dev" {
		err := godotenv.Load(basepath + "/../../../../.env")
		if err != nil {
			log.Printf("Error loading .env files")
		}
	}
}

func ConnectMongoDB() (*mgo.Database, error) {
	var db *mgo.Database

	session, err := mgo.Dial(utils.GetEnv("DB_URI", "localhost"))
	if err != nil {
		return nil, err
	}
	db = session.DB(utils.GetEnv("DB_NAME", "time_record_service"))

	migrate.SetDatabase(db)
	if err := migrate.Up(migrate.AllAvailable); err != nil {
		return nil, err
	}

	return db, nil
}
