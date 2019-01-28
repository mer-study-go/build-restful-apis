package user

import (
	"os"
	"reflect"
	"testing"

	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

func TestMain(m *testing.M) {
	m.Run()
	os.Remove(dbPath)
}

func TestCRUD(t *testing.T) {
	t.Log("Create")
	u := &User{
		ID:   bson.NewObjectId(),
		Name: "John",
		Role: "Tester",
	}
	err := u.Save()
	if err != nil {
		t.Fatalf("Error saving a record: %s", err)
	}

	t.Log("Read")
	u2, err := One(u.ID)
	if err != nil {
		t.Fatalf("Error retrieving a record: %s", err)
	}

	if !reflect.DeepEqual(u2, u) {
		t.Error("Recoreds do not match")
	}

	t.Log("Update")
	u.Role = "developer"
	err = u.Save()
	if err != nil {
		t.Fatalf("Error saving a record: %s", err)
	}
	u3, err := One(u.ID)
	if err != nil {
		t.Fatalf("Error retrieving a record: %s", err)
	}

	if !reflect.DeepEqual(u3, u) {
		t.Error("Records do not match")
	}

	t.Log("Delete")
	err = Delete(u.ID)
	if err != nil {
		t.Fatalf("Error removing a record: %s", err)
	}
	_, err = One(u.ID)
	if err == nil {
		t.Fatal("Record should not exist anymore")
	}

	if err != storm.ErrNotFound {
		t.Fatalf("Error retrieving non-existing record: %s", err)
	}

	t.Log("Read All")
	u2.ID = bson.NewObjectId()
	u3.ID = bson.NewObjectId()

	err = u2.Save()
	if err != nil {
		t.Fatalf("Error saving a record: %s", err)
	}

	err = u3.Save()
	if err != nil {
		t.Fatalf("Error saving a record: %s", err)
	}

	users, err := All()
	if err != nil {
		t.Fatalf("Error reading all record: %s", err)
	}

	if len(users) != 2 {
		t.Errorf("Different number of records retrieved. Expected: 2 / Actual: %d", len(users))
	}

}
