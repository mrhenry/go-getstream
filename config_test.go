package getstream_test

import (
	"os"
	"testing"
	"fmt"
	"errors"
	"time"
	"net/url"

	getstream "github.com/GetStream/stream-go"
)

func TestWerckerConfig(t *testing.T) {

	testAPIKey := os.Getenv("key")
	testAPISecret := os.Getenv("secret")
	testAppID := os.Getenv("app_id")
	testRegion := os.Getenv("region")

	if testAPIKey == "" || testAPISecret == "" || testAppID == "" || testRegion == "" {
		t.Fail()
		return
	}
}

func TestConfig_SetAPIKey(t *testing.T) {
	cfg := getstream.Config{APIKey:"123"}
	if cfg.APIKey != "123" {
		fmt.Println(errors.New(fmt.Sprintf("building cfg didn't set APIKey to '234', got %s", cfg.APIKey)))
		t.Fail()
	}

	chk := cfg.SetAPIKey("234")
	if chk != "234" {
		fmt.Println(errors.New(fmt.Sprintf("SetAPIKey didn't return '234', got %s", chk)))
		t.Fail()
	}
	if cfg.APIKey != "234" {
		fmt.Println(errors.New(fmt.Sprintf("cfg.APIKey isn't 234, got %s", cfg.APIKey)))
		t.Fail()
	}
}

func TestConfig_SetAPISecret(t *testing.T) {
	cfg := getstream.Config{APISecret:"123"}
	if cfg.APISecret != "123" {
		fmt.Println(errors.New(fmt.Sprintf("building cfg didn't set APISecret to '234', got %s", cfg.APISecret)))
		t.Fail()
	}

	chk := cfg.SetAPISecret("234")
	if chk != "234" {
		fmt.Println(errors.New(fmt.Sprintf("SetAPISecret didn't return '234', got %s", chk)))
		t.Fail()
	}
	if cfg.APISecret != "234" {
		fmt.Println(errors.New(fmt.Sprintf("cfg.APISecret isn't 234, got %s", cfg.APISecret)))
		t.Fail()
	}
}

func TestConfig_SetAppID(t *testing.T) {
	cfg := getstream.Config{AppID:"123"}
	if cfg.AppID != "123" {
		fmt.Println(errors.New(fmt.Sprintf("building cfg didn't set AppID to '234', got %s", cfg.AppID)))
		t.Fail()
	}

	chk := cfg.SetAppID("234")
	if chk != "234" {
		fmt.Println(errors.New(fmt.Sprintf("SetAppID didn't return '234', got %s", chk)))
		t.Fail()
	}
	if cfg.AppID != "234" {
		fmt.Println(errors.New(fmt.Sprintf("cfg.AppID isn't 234, got %s", cfg.AppID)))
		t.Fail()
	}
}

func TestConfig_SetLocation(t *testing.T) {
	cfg := getstream.Config{Location:"123"}
	if cfg.Location != "123" {
		fmt.Println(errors.New(fmt.Sprintf("building cfg didn't set Location to '234', got %s", cfg.Location)))
		t.Fail()
	}

	chk := cfg.SetLocation("234")
	if chk != "234" {
		fmt.Println(errors.New(fmt.Sprintf("SetLocation didn't return '234', got %s", chk)))
		t.Fail()
	}
	if cfg.Location != "234" {
		fmt.Println(errors.New(fmt.Sprintf("cfg.Location isn't 234, got %s", cfg.Location)))
		t.Fail()
	}
}

func TestConfig_SetTimeout(t *testing.T) {
	cfg := getstream.Config{TimeoutInt:123, TimeoutDuration:time.Duration(456)}
	if cfg.TimeoutInt != 123 {
		fmt.Println(errors.New(fmt.Sprintf("building cfg didn't set TimeoutInt to '123', got %d", cfg.TimeoutInt)))
		t.Fail()
	}
	if cfg.TimeoutDuration != 456 {
		fmt.Println(errors.New(fmt.Sprintf("building cfg didn't set timeoutDuration to time.Duration(456), got %d", cfg.TimeoutDuration)))
		t.Fail()
	}

	chk := cfg.SetTimeout(234)
	if chk != time.Duration(234 * time.Second) {
		fmt.Println(errors.New(fmt.Sprintf("SetTimeoutInt didn't return time.Duration(234), got %d", chk)))
		t.Fail()
	}
	if cfg.TimeoutInt != 234 {
		fmt.Println(errors.New(fmt.Sprintf("cfg.TimeoutInt isn't int(234), got %s", cfg.TimeoutInt)))
		t.Fail()
	}
	if cfg.TimeoutDuration != time.Duration(234 * time.Second) {
		fmt.Println(errors.New(fmt.Sprintf("setting cfg.TimeoutInt didn't set timeoutDuraction to time.Duraction(234), got %d", cfg.TimeoutInt)))
		t.Fail()
	}
}

func TestConfig_SetVersion(t *testing.T) {
	cfg := getstream.Config{Version:"123"}
	if cfg.Version != "123" {
		fmt.Println(errors.New(fmt.Sprintf("building cfg didn't set Version to '234', got %s", cfg.Version)))
		t.Fail()
	}

	chk := cfg.SetVersion("234")
	if chk != "234" {
		fmt.Println(errors.New(fmt.Sprintf("SetVersion didn't return '234', got %s", chk)))
		t.Fail()
	}
	if cfg.Version != "234" {
		fmt.Println(errors.New(fmt.Sprintf("cfg.Version isn't 234, got %s", cfg.Version)))
		t.Fail()
	}
}

func TestConfig_SetBaseURL(t *testing.T) {
	url, _ := url.Parse("http://api.getstream.io")
	cfg := getstream.Config{BaseURL:url}
	if cfg.BaseURL != url {
		fmt.Println(errors.New(fmt.Sprintf("building cfg didn't set BaseURL to 'http://use-east-api.getstream.io', got %s", cfg.BaseURL)))
		t.Fail()
	}

	url, _ = url.Parse("http://use-east-api.getstream.io")
	chk := cfg.SetBaseURL(url)
	if chk != url {
		fmt.Println(errors.New(fmt.Sprintf("SetBaseURL didn't return 'http://use-east-api.getstream.io', got %s", chk)))
		t.Fail()
	}
	if cfg.BaseURL != url {
		fmt.Println(errors.New(fmt.Sprintf("cfg.BaseURL isn't http://use-east-api.getstream.io, got %s", cfg.BaseURL)))
		t.Fail()
	}
}