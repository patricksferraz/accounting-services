package external

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
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

func ConnectAuthService() (*grpc.ClientConn, error) {

	conn, err := grpc.Dial(os.Getenv("AUTH_SERVICE_ADDR"), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
