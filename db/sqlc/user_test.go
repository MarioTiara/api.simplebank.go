package db

import (
	"context"
	"testing"
	"time"

	"github.com/marioTiara/api.simplebank.go/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       utils.RandomOwner(),
		HashedPassword: "secret",
		FullName:       utils.RandomOwner(),
		Email:          utils.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	createdUser := createRandomUser(t)
	userfromDB, err := testQueries.GetUser(context.Background(), createdUser.Username)

	require.NoError(t, err)
	require.NotEmpty(t, userfromDB)

	require.Equal(t, createdUser.Username, userfromDB.Username)
	require.Equal(t, createdUser.HashedPassword, userfromDB.HashedPassword)
	require.Equal(t, createdUser.FullName, userfromDB.FullName)
	require.Equal(t, createdUser.Email, userfromDB.Email)
	require.WithinDuration(t, userfromDB.CreatedAt, createdUser.CreatedAt, time.Second)
}
