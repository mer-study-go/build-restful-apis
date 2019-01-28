package handlers

import (
	"net/http"
	"github.com/build-restful-apis/src/user"
	"testing"
	"reflect"
	"io/ioutil"
	"bytes"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

func TestBodyToUser(t *testing.T) {
	valid := &user.User {
		ID: bson.NewObjectId(),
		Name: "John",
		Role: "Tester",
	}

	js, err := json.Marshal(valid)

	if err != nil {
		t.Errorf("Error marshalling a valid users: %s", err)
		t.FailNow()

	}

	ts := []struct {
		txt string
		r *http.Request
		u *user.User
		err bool
		exp *user.User
	} {
		// request is nil
		{
			txt : "nil request",
			err: true,
		},
		// request is empty
		{
			txt: "empty request body",
			r: &http.Request{},
			err: true,
		},
		// empty user
		{
			txt: "emtpy user", 
			r: &http.Request {
				Body: ioutil.NopCloser(bytes.NewBufferString("{}")),
			},
			err: true,
		},
		// malformed data
		{
			txt: "malformed data",
			r: &http.Request {
				Body: ioutil.NopCloser(bytes.NewBufferString(`{"id":12}`)),
			},
			u: &user.User{},
			err: true,
		},
		// valid data 
		{
			txt: "valid request", 
			r: &http.Request {
				Body: ioutil.NopCloser(bytes.NewBuffer(js)),
			},
			u: &user.User{},
			exp: valid,
		},
	}

	for _, tc := range ts {
		t.Log(tc.txt)
		err := bodyToUser(tc.r, tc.u)
		if tc.err {
			if err == nil {
				t.Error("Expected error, got none.")
			}
			continue
		}

		if err != nil {
			t.Errorf("Unexpected error: %s", err)
			continue
		}

		if !reflect.DeepEqual(tc.u, tc.exp) {
			t.Error("Unmarsshaled data is different:")
			t.Error(tc.u)
			t.Error(tc.exp)
		}
	}
	
}