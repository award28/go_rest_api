package sessionStore_test

import (
	"go_rest_api/internal"
	"go_rest_api/internal/sessionStore"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	testUsername = "name"
	testEmail    = "a@test.com"
	testPassword = "pass"
)

func Test_User_Store(t *testing.T) {
	t.Run("Create and Retrieve session user", create_and_retrieve_session_user)
	t.Run("Verify session user deletion", verify_session_user_deletion)
}

func user_setup() (*sessionStore.UserStore, http.ResponseWriter, *http.Request) {
	store := sessionStore.NewStore(storeName)
	userStore := sessionStore.NewUserStore(store)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	return userStore, w, r
}

func create_and_retrieve_session_user(t *testing.T) {
	store, w, r := user_setup()
	defer store.DeleteSessionUser(r, w)

	log_error = func(t string, err error) {
		log.Fatalf("Unable to %s session user, reason: %s", t, err.Error())
	}

	actualUser := &root.User{testUsername, testEmail, testPassword}
	err := store.SetSessionUser(r, w, actualUser)
	if err != nil {
		log_set_error("set", err)
	}

	responseUser, err := store.GetSessionUser(r)
	if err != nil {
		log_set_error("get", err)
	}

	if actualUser != responseUser {
		log.Fatalf("The user stored differs from the original.")
	}
}
func verify_session_user_deletion(t *testing.T) {
	store, w, r := user_setup()

	err := store.SetSessionUser(r, w, &root.User{testUsername, testEmail, testPassword})
	if err != nil {
		log_set_error("set", err)
	}

	store.DeleteSessionUser(r, w)

	_, err = store.GetSessionUser(r)
	if err == nil {
		log.Fatalf("Was able to retrieve session user after deletion.")
	}
}
