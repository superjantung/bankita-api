package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/superjantung/bankita-api/util"
)

func assertUserCreated(t *testing.T, user User, arg CreateUserParams) {
	require.NotZero(t, user.Username)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)
}

func assertUserEqual(t *testing.T, expected, actual User) {
	require.Equal(t, expected.Username, actual.Username)
	require.Equal(t, expected.HashedPassword, actual.HashedPassword)
	require.Equal(t, expected.FullName, actual.FullName)
	require.Equal(t, expected.Email, actual.Email)
	require.WithinDuration(t, expected.PasswordChangedAt, actual.PasswordChangedAt, time.Second)
	require.WithinDuration(t, expected.CreatedAt, actual.CreatedAt, time.Second)
}

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: "secret",
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	createdUser, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, createdUser)

	assertUserCreated(t, createdUser, arg)
	return createdUser
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	createdUser := createRandomUser(t)
	retrievedUser, err := testQueries.GetUser(context.Background(), createdUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedUser)

	assertUserEqual(t, createdUser, retrievedUser)
}
