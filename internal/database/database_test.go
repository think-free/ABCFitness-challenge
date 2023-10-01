package database_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/think-free/ABCFitness-challenge/internal/cliparams"
	"github.com/think-free/ABCFitness-challenge/internal/database"
	"github.com/think-free/ABCFitness-challenge/internal/datamodel"
)

// Testing the implementation
func TestSaveUser(t *testing.T) {
	ctx := context.Background()
	db := getDatabase(ctx)

	user := &datamodel.User{
		ID: "1",
		BaseUser: datamodel.BaseUser{
			Name:    "Elon",
			Surname: "Musk",
			Email:   "elon.musk@example.com",
			Phone:   "+341234567890",
		},
	}

	err := db.SaveUser(ctx, user)
	assert.NoError(t, err)

	// Saving the same user again should return an error
	err = db.SaveUser(ctx, user)
	assert.Error(t, err)
}

func TestGetUserByID(t *testing.T) {
	ctx := context.Background()
	db := getDatabase(ctx)

	user := &datamodel.User{
		ID: "1",
		BaseUser: datamodel.BaseUser{
			Name:    "Elon",
			Surname: "Musk",
			Email:   "elon.musk@example.com",
			Phone:   "+341234567890",
		},
	}

	err := db.SaveUser(ctx, user)
	assert.NoError(t, err)

	// Retrieving the user by ID should return the same user
	retrievedUser, err := db.GetUserByID(ctx, user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user, retrievedUser)

	// Retrieving a non-existent user should return an error
	_, err = db.GetUserByID(ctx, "2")
	assert.Error(t, err)
}

func TestGetUserID(t *testing.T) {
	ctx := context.Background()
	db := getDatabase(ctx)

	user := &datamodel.User{
		ID: "1",
		BaseUser: datamodel.BaseUser{
			Name:    "Elon",
			Surname: "Musk",
			Email:   "elon.musk@example.com",
			Phone:   "+341234567890",
		},
	}

	err := db.SaveUser(ctx, user)
	assert.NoError(t, err)

	// Retrieving the ID of an existing user should return the user's ID
	id, err := db.GetUserID(ctx, user)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, id)

	// Retrieving the ID of a non-existent user should return an error
	_, err = db.GetUserID(ctx, &datamodel.User{
		BaseUser: datamodel.BaseUser{
			Name:    "Leah",
			Surname: "Mars",
			Email:   "leah.mars@example.com",
			Phone:   "+341234567890",
		},
	})
	assert.Error(t, err)
}

func TestListUsers(t *testing.T) {
	ctx := context.Background()
	db := getDatabase(ctx)

	user1 := &datamodel.User{
		ID: "1",
		BaseUser: datamodel.BaseUser{
			Name:    "Elon",
			Surname: "Musk",
			Email:   "elon.musk@example.com",
			Phone:   "+341234567890",
		},
	}

	user2 := &datamodel.User{
		ID: "2",
		BaseUser: datamodel.BaseUser{
			Name:    "Leah",
			Surname: "Mars",
			Email:   "leah.mars@example.com",
			Phone:   "+341234567890",
		},
	}

	err := db.SaveUser(ctx, user1)
	assert.NoError(t, err)

	err = db.SaveUser(ctx, user2)
	assert.NoError(t, err)

	// Listing users should return all saved users
	users, err := db.ListUsers(ctx, 0, 0)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []*datamodel.User{user1, user2}, users)
}

func TestSaveClass(t *testing.T) {
	ctx := context.Background()
	db := getDatabase(ctx)

	start := time.Now().AddDate(0, 0, -20)
	end := time.Now().AddDate(0, 0, -10)

	class := &datamodel.Class{
		ID: "1",
		BaseClass: datamodel.BaseClass{
			Studio:        "Yoga Studio",
			Name:          "Yoga Class",
			StartDate:     &start,
			EndDate:       &end,
			DailyCapacity: 20,
		},
	}

	err := db.SaveClass(ctx, class)
	assert.NoError(t, err)

	// Saving the same class again should return an error
	err = db.SaveClass(ctx, class)
	assert.Error(t, err)
}

func getDatabase(ctx context.Context) database.Database {
	return database.New(ctx, cliparams.New())
}
