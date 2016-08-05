package getstream

import (
	"os"
	"testing"
)

func TestWerckerConfig(t *testing.T) {

	testAPIKey := os.Getenv("key")
	testAPISecret := os.Getenv("secret")
	testAppID := os.Getenv("app_id")
	testRegion := os.Getenv("region")

	if testAPIKey == "" || testAPISecret == "" || testAppID == "" || testRegion == "" {
		t.Fail()
	}
}

func TestFail(t *testing.T) {
	t.Fail()
}
