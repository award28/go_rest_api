package sessionStore_test

import (
	"encoding/gob"
	"errors"
	"go_rest_api/pkg/sessionStore"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	storeName = "Test_Store"
	valueName = "test"
)

type testStruct struct {
	Name string
}

func Test_Store(t *testing.T) {
	t.Run("Create and Retrieve session value", create_and_retrieve_session_value)
	t.Run("Create and Retrieve session struct", create_and_retrieve_session_struct)
}

func setup() (*sessionStore.Store, http.ResponseWriter, *http.Request) {
	store := sessionStore.NewStore(storeName)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	return store, w, r
}

func create_and_retrieve_session_value(t *testing.T) {
	store, w, r := setup()
	defer store.DeleteStore()

	actualValue := "actual value"
	err := store.Set(r, w, func(s *sessionStore.Session) error {
		s.Values[valueName] = actualValue
		return nil
	})
	if err != nil {
		log.Fatalf("Unable to set value `"+valueName+"`, reason: %s", err.Error())
	}

	var responseValue string
	err = store.Get(r, func(s *sessionStore.Session) error {
		var ok bool
		responseValue, ok = s.Values[valueName].(string)
		if !ok {
			return errors.New("Value `" + valueName + "` is not in the current session")
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Unable to get value `"+valueName+"`, reason: %s", err.Error())
	}

	if actualValue != responseValue {
		log.Fatalf("The value stored differs from the original.")
	}
}

func create_and_retrieve_session_struct(t *testing.T) {
	gob.Register(&testStruct{})

	store, w, r := setup()
	defer store.DeleteStore()

	actualValue := &testStruct{valueName}
	err := store.Set(r, w, func(s *sessionStore.Session) error {
		s.Values[valueName] = actualValue
		return nil
	})
	if err != nil {
		log.Fatalf("Unable to set struct value, reason: %s", err.Error())
	}

	responseValue := new(testStruct)

	err = store.Get(r, func(s *sessionStore.Session) error {
		val, ok := s.Values[valueName]
		if !ok {
			return errors.New("Value `" + valueName + "` is not in the current session")
		}

		responseValue, ok = val.(*testStruct)
		if !ok {
			return errors.New("Value `" + valueName + "` is not of type testStruct")
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Unable to get struct value, reason: %s", err.Error())
	}

	if actualValue != responseValue {
		log.Fatalf("The struct value stored differs from the original.")
	}
}
