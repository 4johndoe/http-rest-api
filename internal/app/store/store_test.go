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
		databaseURL = "host=10.133.133.26 port=5555 dbname=restapi_test sslmode=disable user=crm password=crm26"
	}

	os.Exit(m.Run())
}
