package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/superjantung/bankita-api/util"
)

func assertTransferCreated(t *testing.T, transfer Transfer, arg CreateTransferParams) {
	require.NotZero(t, transfer.ID)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.NotZero(t, transfer.CreatedAt)
}

func assertTransferEqual(t *testing.T, expected, actual Transfer) {
	require.Equal(t, expected.ID, actual.ID)
	require.Equal(t, expected.FromAccountID, actual.FromAccountID)
	require.Equal(t, expected.ToAccountID, actual.ToAccountID)
	require.Equal(t, expected.Amount, actual.Amount)
	require.WithinDuration(t, expected.CreatedAt, actual.CreatedAt, time.Second)
}

func createRandomTransfer(t *testing.T, fromAccount, toAccount Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        util.RandomBalance(),
	}

	createdTransfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, createdTransfer)

	assertTransferCreated(t, createdTransfer, arg)
	return createdTransfer
}

func TestCreateTransfer(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	createRandomTransfer(t, fromAccount, toAccount)
}

func TestGetTransfer(t *testing.T) {
	createdFromAccount := createRandomAccount(t)
	createdToAccount := createRandomAccount(t)
	createdTransfer := createRandomTransfer(t, createdFromAccount, createdToAccount)

	retrievedTransfer, err := testQueries.GetTransfer(context.Background(), createdTransfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedTransfer)

	assertTransferEqual(t, createdTransfer, retrievedTransfer)
}

func TestListTransfers(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	// Create multiple random transfers
	numTransfers := 10
	createdTransfers := make([]Transfer, numTransfers)
	for i := 0; i < numTransfers; i++ {
		createdTransfers[i] = createRandomTransfer(t, fromAccount, toAccount)
		createdTransfers[i] = createRandomTransfer(t, toAccount, fromAccount)
	}

	// Retrieve a list of transfers
	arg := ListTransfersParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Limit:         5,
		Offset:        5,
	}
	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	// Assert that the retrieved transfers are not empty
	for _, retrievedTransfer := range transfers {
		require.NotEmpty(t, retrievedTransfer)
		require.True(t, retrievedTransfer.FromAccountID == fromAccount.ID || retrievedTransfer.ToAccountID == fromAccount.ID)
	}
}
