package db

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	mgo "gopkg.in/mgo.v2"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}
}

func ConnectDB() *mgo.Database {
	var db *mgo.Database

	session, err := mgo.Dial(os.Getenv("DB_URI"))
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(os.Getenv("DB_NAME"))

	return db
}
