package bolt_test

import (
	"go_rest_api/pkg"
	"go_rest_api/pkg/bolt"
	"go_rest_api/pkg/mock"
	"log"
	"testing"
)

const (
	dbName         = "go_rest_api.db"
	userBucketName = "test_users"
)

func Test_UserService(t *testing.T) {
	t.Run("Create And Retrieve User", insert_and_retrieve_user_from_bolt)
}

func insert_and_retrieve_user_from_bolt(t *testing.T) {
	//Arrange
	db, err := bolt.NewDatabase()
	if err != nil {
		log.Fatalf("Unable to connect to bolt: %s", err)
	}
	defer func() {
		db.DeleteBucket(userBucketName)
		db.Close()
	}()

	userService := bolt.NewUserService(db, userBucketName, &mock.Hash{})

	testUsername := "integration_test_user"
	testPassword := "integration_test_password"
	user := root.User{
		Username: testUsername,
		Password: testPassword}

	//Act
	err = userService.Create(&user)

	//Assert
	if err != nil {
		t.Errorf("Unable to create user: %s", err)
	}
	result, err := userService.GetByUsername(testUsername)
	if err != nil {
		t.Errorf("Could not get user: %s", err)
	}

	if result.Username != user.Username {
		t.Errorf("Incorrect Username. Expected `%s`, Got: `%s`", testUsername, result.Username)
	} else if result.Password != user.Password {
		t.Errorf("Incorrect Password. Expected `%s`, Got: `%s`", testPassword, result.Password)
	}
}
