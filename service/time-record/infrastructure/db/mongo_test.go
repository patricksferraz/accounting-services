package db_test

import (
	"os"
	"testing"

	"github.com/patricksferraz/accounting-services/service/time-record/infrastructure/db"
	"github.com/stretchr/testify/require"
)

func TestDb_ConnectMongoDB(t *testing.T) {

	database, err := db.ConnectMongoDB()
	require.Nil(t, err)
	require.NotEmpty(t, database)

	os.Setenv("DB_URI", "mongodb://1.1.1.1")
	_, err = db.ConnectMongoDB()
	require.NotNil(t, err)
}
