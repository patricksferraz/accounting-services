package db

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/patricksferraz/accounting-services/utils"
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

	return db, nil
}
