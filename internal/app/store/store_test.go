package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=10.133.133.26 port=5555 user=crm password=crm1 dbname=restapi_test sslmode=disable"
	}

	os.Exit(m.Run())
}