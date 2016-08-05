package getstream

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestError(t *testing.T) {

	errorResponse := "{\"code\": 5, \"detail\": \"some detail\", \"duration\": \"36ms\", \"exception\": \"an exception\", \"status_code\": 400}"

	var getStreamError Error
	err := json.Unmarshal([]byte(errorResponse), &getStreamError)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	testError := Error{
		Code:        5,
		Detail:      "some detail",
		RawDuration: "36ms",
		Exception:   "an exception",
		StatusCode:  400,
	}

	if getStreamError != testError {
		fmt.Println(err)
		t.Fail()
	}

	if getStreamError.Duration() != time.Millisecond*36 {
		fmt.Println(err)
		t.Fail()
	}

	if getStreamError.Error() != "an exception (36ms): some detail" {
		fmt.Println(err)
		t.Fail()
	}
}

func TestErrorBadDuration(t *testing.T) {

	testError := Error{
		Code:        5,
		Detail:      "some detail",
		RawDuration: "36blah",
		Exception:   "an exception",
		StatusCode:  400,
	}

	if testError.Duration() != time.Duration(0) {
		fmt.Println(err)
		t.Fail()
		return
	}

}
