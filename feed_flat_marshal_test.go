package getstream

import "github.com/pborman/uuid"
import "time"
import "testing"
import "encoding/json"
import "fmt"

func TestActivityMetaData(t *testing.T) {

	fmt.Println("Meta Data : ")

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

	activity := FlatFeedActivity{
		ForeignID: uuid.New(),
		Actor:     FeedID("user:eric"),
		Object:    FeedID("user:bob"),
		Target:    FeedID("user:john"),
		Verb:      "post",
		TimeStamp: &now,
		Data:      dataB,
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

	fmt.Println(string(b))

	resultActivity := FlatFeedActivity{}
	err = json.Unmarshal(b, &resultActivity)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	resultActivity2 := FlatFeedActivity{}
	err = json.Unmarshal(b2, &resultActivity2)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if resultActivity.ForeignID != activity.ForeignID {
		fmt.Println(resultActivity.ForeignID)
		fmt.Println(activity.ForeignID)
		t.Fail()
	}
	if resultActivity.Actor != activity.Actor {
		fmt.Println(resultActivity.Actor)
		fmt.Println(activity.Actor)
		t.Fail()
	}
	if resultActivity.Verb != activity.Verb {
		fmt.Println(resultActivity.Verb)
		fmt.Println(activity.Verb)
		t.Fail()
	}
	if resultActivity.Object != activity.Object {
		fmt.Println(resultActivity.Object)
		fmt.Println(activity.Object)
		t.Fail()
	}
	if resultActivity.Target != activity.Target {
		fmt.Println(resultActivity.Target)
		fmt.Println(activity.Target)
		t.Fail()
	}
	if resultActivity.TimeStamp != activity.TimeStamp {
		fmt.Println(resultActivity.TimeStamp)
		fmt.Println(activity.TimeStamp)
		t.Fail()
	}
	if resultActivity.MetaData["meta"] != activity.MetaData["meta"] {
		fmt.Println(resultActivity.MetaData)
		fmt.Println(activity.MetaData)
		t.Fail()
	}
	if string(resultActivity.Data) != string(activity.Data) {
		fmt.Println(string(resultActivity.Data))
		fmt.Println(string(activity.Data))
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
