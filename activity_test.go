package getstream

import (
	"encoding/json"
	"reflect"
	"time"

	"github.com/pborman/uuid"
)

import "testing"

import "fmt"

func TestActivityMarshalling(t *testing.T) {

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

	activity := Activity{
		ForeignID: uuid.New(),
		Actor:     FeedID("user:eric"),
		Object:    FeedID("user:bob"),
		Target:    FeedID("user:john"),
		Verb:      "post",
		TimeStamp: &now,
		Data:      &raw,
		MetaData: map[string]interface{}{
			"stringKey": "stringValue",
			"intKey":    1,
			"floatKey":  1.235,
			"boolKey":   true,
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

	resultActivity := Activity{}
	err = json.Unmarshal(b, &resultActivity)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	resultActivity2 := Activity{}
	err = json.Unmarshal(b2, &resultActivity2)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	if resultActivity.ForeignID != activity.ForeignID {
		fmt.Println(activity.ForeignID)
		fmt.Println(resultActivity.ForeignID)
		t.Fail()
		return
	}
	if resultActivity.Actor != activity.Actor {
		fmt.Println(activity.Actor)
		fmt.Println(resultActivity.Actor)
		t.Fail()
		return
	}
	if resultActivity.Verb != activity.Verb {
		fmt.Println(activity.Verb)
		fmt.Println(resultActivity.Verb)
		t.Fail()
		return
	}
	if resultActivity.Object != activity.Object {
		fmt.Println(activity.Object)
		fmt.Println(resultActivity.Object)
		t.Fail()
		return
	}
	if resultActivity.Target != activity.Target {
		fmt.Println(activity.Target)
		fmt.Println(resultActivity.Target)
		t.Fail()
		return
	}
	if resultActivity.TimeStamp.Format("2006-01-02T15:04:05.999999") != activity.TimeStamp.Format("2006-01-02T15:04:05.999999") {
		fmt.Println(activity.TimeStamp)
		fmt.Println(resultActivity.TimeStamp)
		t.Fail()
		return
	}

	if string(*resultActivity.Data) != string(*activity.Data) {
		fmt.Println(string(*activity.Data))
		fmt.Println(string(*resultActivity.Data))
		t.Fail()
		return
	}

	vString, okString := resultActivity.MetaData["stringKey"].(string)
	if !okString {
		fmt.Println(reflect.TypeOf(resultActivity.MetaData["stringKey"]))
		fmt.Println("Not a String")
		t.Fail()
		return
	}
	if vString != "stringValue" {
		fmt.Println("Not the correct value")
		t.Fail()
		return
	}

	vInt, okInt := resultActivity.MetaData["intKey"].(int)
	if !okInt {
		fmt.Println(reflect.TypeOf(resultActivity.MetaData["intKey"]))
		fmt.Println("Not an Int")
		t.Fail()
		return
	}
	if vInt != 1 {
		fmt.Println("Not the correct value")
		t.Fail()
		return
	}

	vFloat, okFloat := resultActivity.MetaData["floatKey"].(float64)
	if !okFloat {
		fmt.Println(reflect.TypeOf(resultActivity.MetaData["floatKey"]))
		fmt.Println("Not a Float")
		t.Fail()
		return
	}
	if vFloat != 1.235 {
		fmt.Println("Not the correct value")
		t.Fail()
		return
	}

	vBool, okBool := resultActivity.MetaData["boolKey"].(bool)
	if !okBool {
		fmt.Println(reflect.TypeOf(resultActivity.MetaData["boolKey"]))
		fmt.Println("Not a Bool")
		t.Fail()
		return
	}
	if vBool != true {
		fmt.Println("Not the correct value")
		t.Fail()
		return
	}

}
