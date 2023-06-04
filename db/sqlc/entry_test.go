package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/superjantung/bankita-api/util"
)

func assertEntryCreated(t *testing.T, entry Entry, arg CreateEntryParams) {
	require.NotZero(t, entry.ID)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.CreatedAt)
}

func assertEntryEqual(t *testing.T, expected, actual Entry) {
	require.Equal(t, expected.ID, actual.ID)
	require.Equal(t, expected.AccountID, actual.AccountID)
	require.Equal(t, expected.Amount, actual.Amount)
	require.WithinDuration(t, expected.CreatedAt, actual.CreatedAt, time.Second)
}

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomBalance(),
	}

	createdEntry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, createdEntry)

	assertEntryCreated(t, createdEntry, arg)
	return createdEntry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	createdAccount := createRandomAccount(t)
	createdEntry := createRandomEntry(t, createdAccount)
	retrievedEntry, err := testQueries.GetEntry(context.Background(), createdEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedEntry)

	assertEntryEqual(t, createdEntry, retrievedEntry)
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)

	// Create multiple random entries
	numEntries := 10
	createdEntries := make([]Entry, numEntries)
	for i := 0; i < numEntries; i++ {
		createdEntries[i] = createRandomEntry(t, account)
	}

	// Retrieve a list of entries
	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	// Assert that the retrieved entries are not empty
	for _, retrievedEntry := range entries {
		require.NotEmpty(t, retrievedEntry)
		require.Equal(t, arg.AccountID, retrievedEntry.AccountID)
	}
}
