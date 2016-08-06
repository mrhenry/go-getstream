package getstream

import "github.com/pborman/uuid"
import "time"
import "testing"
import "encoding/json"

func TestActivityMetaData(t *testing.T) {

	now := time.Now()

	data := struct {
		Foo  string
		Fooz string
	}{
		Foo:  "foo",
		Fooz: "fooz",
	}

	dataB, err := json.Marshal(data)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	activity := FlatFeedActivity{
		ForeignID: uuid.New(),
		Actor:     FeedID("user:eric"),
		Object:    FeedID("user:bob"),
		Verb:      "post",
		TimeStamp: &now,
		Data:      dataB,
		MetaData: map[string]string{
			"meta": "data",
		},
	}

	b, err := prepareForGetstream(&activity)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	t.Log(string(b))

	output := extractFromGetStream(b)
	resultActivity := output.activity()

	t.Log(resultActivity)

}
