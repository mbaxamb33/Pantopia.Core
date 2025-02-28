package db

import (
	"context"
	"testing"

	"github.com/mbaxamb33/pantopia/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) Users {
	account := createRandomAccount(t) // Ensure account exists before creating a user

	arg := CreateUserParams{
		AccountID: account.ID,
		Email:     util.RandomEmail(),
		FullName:  util.RandomName(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.AccountID, user.AccountID)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.FullName, user.FullName)

	require.NotZero(t, user.ID)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.AccountID, user2.AccountID)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.FullName, user2.FullName)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)
	arg := UpdateUserParams{
		ID:       user1.ID,
		FullName: util.RandomName(),
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.AccountID, user2.AccountID)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, arg.FullName, user2.FullName)
}

func TestDeleteUser(t *testing.T) {
	user := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)

	_, err = testQueries.GetUser(context.Background(), user.ID)
	require.Error(t, err)
}

func TestListUsers(t *testing.T) {
	account := createRandomAccount(t) // Create a single account

	for i := 0; i < 5; i++ {
		arg := CreateUserParams{
			AccountID: account.ID, // Use the same account
			Email:     util.RandomEmail(),
			FullName:  util.RandomName(),
		}
		_, err := testQueries.CreateUser(context.Background(), arg)
		require.NoError(t, err)
	}

	arg := ListUsersParams{
		AccountID: account.ID, // Use the correct account
		Limit:     5,
		Offset:    0,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5) // Now expects exactly 5 users

	for _, user := range users {
		require.NotEmpty(t, user)
		require.Equal(t, account.ID, user.AccountID) // Ensure they all belong to the same account
	}
}
