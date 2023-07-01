package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/superjantung/bankita-api/util"
)

func assertAccountCreated(t *testing.T, account Account, arg CreateAccountParams) {
	require.NotZero(t, account.ID)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.CreatedAt)
}

func assertAccountEqual(t *testing.T, expected, actual Account) {
	require.Equal(t, expected.ID, actual.ID)
	require.Equal(t, expected.Owner, actual.Owner)
	require.Equal(t, expected.Balance, actual.Balance)
	require.Equal(t, expected.Currency, actual.Currency)
	require.WithinDuration(t, expected.CreatedAt, actual.CreatedAt, time.Second)
}

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrencies(),
	}

	createdAccount, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, createdAccount)

	assertAccountCreated(t, createdAccount, arg)
	return createdAccount
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)
	retrievedAccount, err := testQueries.GetAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedAccount)

	assertAccountEqual(t, createdAccount, retrievedAccount)
}

func TestUpdateAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      createdAccount.ID,
		Balance: util.RandomBalance(),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	// Retrieve the updated account to compare the changes
	retrievedAccount, err := testQueries.GetAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedAccount)

	assertAccountEqual(t, updatedAccount, retrievedAccount)

	// Assert that the balance of the updated account matches the provided balance in the arg parameter
	require.Equal(t, arg.Balance, updatedAccount.Balance)
}

func TestDeleteAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)

	// Try to retrieve the deleted account
	deletedAccount, err := testQueries.GetAccount(context.Background(), createdAccount.ID)
	require.Error(t, err)

	// Assert that the error is of type sql.ErrNoRows
	require.Equal(t, sql.ErrNoRows, err)

	// Assert that the deleted account is empty
	require.Empty(t, deletedAccount)
}

func TestListAccounts(t *testing.T) {
	var lastAccount Account

	// Create multiple random accounts
	numAccounts := 10
	for i := 0; i < numAccounts; i++ {
		lastAccount = createRandomAccount(t)
	}

	// Retrieve a list of accounts
	arg := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	// Assert that the retrieved accounts are not empty
	for _, retrievedAccount := range accounts {
		require.NotEmpty(t, retrievedAccount)
		require.Equal(t, lastAccount.Owner, retrievedAccount.Owner)
	}
}
