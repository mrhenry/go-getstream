package getstream_test

import "github.com/pborman/uuid"
import "time"
import "testing"
import "encoding/json"
import "fmt"
import 	getstream "github.com/GetStream/stream-go"


func TestAggregatedActivityMetaData(t *testing.T) {

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
		fmt.Println(err)
		t.Fail()
		return
	}

	raw := json.RawMessage(dataB)

	activity := getstream.Activity{
		ForeignID: uuid.New(),
		Actor:     getstream.FeedID("user:eric"),
		Object:    getstream.FeedID("user:bob"),
		Target:    getstream.FeedID("user:john"),
		Origin:    getstream.FeedID("user:barry"),
		Verb:      "post",
		TimeStamp: &now,
		Data:      &raw,
		MetaData: map[string]string{
			"meta": "data",
		},
	}

	b, err := json.Marshal(&activity)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	b2, err := json.Marshal(activity)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	resultActivity := getstream.Activity{}
	err = json.Unmarshal(b, &resultActivity)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	resultActivity2 := getstream.Activity{}
	err = json.Unmarshal(b2, &resultActivity2)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if resultActivity.ForeignID != activity.ForeignID {
		fmt.Println(activity.ForeignID)
		fmt.Println(resultActivity.ForeignID)
		t.Fail()
	}
	if resultActivity.Actor != activity.Actor {
		fmt.Println(activity.Actor)
		fmt.Println(resultActivity.Actor)
		t.Fail()
	}
	if resultActivity.Origin != activity.Origin {
		fmt.Println(activity.Origin)
		fmt.Println(resultActivity.Origin)
		t.Fail()
	}
	if resultActivity.Verb != activity.Verb {
		fmt.Println(activity.Verb)
		fmt.Println(resultActivity.Verb)
		t.Fail()
	}
	if resultActivity.Object != activity.Object {
		fmt.Println(activity.Object)
		fmt.Println(resultActivity.Object)
		t.Fail()
	}
	if resultActivity.Target != activity.Target {
		fmt.Println(activity.Target)
		fmt.Println(resultActivity.Target)
		t.Fail()
	}
	if resultActivity.TimeStamp.Format("2006-01-02T15:04:05.999999") != activity.TimeStamp.Format("2006-01-02T15:04:05.999999") {
		fmt.Println(activity.TimeStamp)
		fmt.Println(resultActivity.TimeStamp)
		t.Fail()
	}
	if resultActivity.MetaData["meta"] != activity.MetaData["meta"] {
		fmt.Println(activity.MetaData)
		fmt.Println(resultActivity.MetaData)
		t.Fail()
	}
	if string(*resultActivity.Data) != string(*activity.Data) {
		fmt.Println(string(*activity.Data))
		fmt.Println(string(*resultActivity.Data))
		t.Fail()
	}

	// fmt.Println(resultActivity)
	// fmt.Println(resultActivity.ForeignID)
	// fmt.Println(string(resultActivity.Data))
	// fmt.Println(resultActivity.MetaData)
	//
	// fmt.Println(resultActivity2)
	// fmt.Println(resultActivity2.ForeignID)
	// fmt.Println(string(resultActivity2.Data))
	// fmt.Println(resultActivity2.MetaData)

}
