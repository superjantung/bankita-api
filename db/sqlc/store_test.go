package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	fmt.Println(">> before:", fromAccount.Balance, toAccount.Balance)

	// Run concurrent transactions
	n := 5
	amount := int64(100000)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			ctx := context.Background()
			result, err := store.TransferTx(ctx, TransferTxParams{
				FromAccountID: fromAccount.ID,
				ToAccountID:   toAccount.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// Check results
	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// Check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, fromAccount.ID, transfer.FromAccountID)
		require.Equal(t, toAccount.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// Check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, fromAccount.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, toAccount.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// Check accounts
		fromAccountResult := result.FromAccount
		require.NotEmpty(t, fromAccountResult)
		require.Equal(t, fromAccountResult.ID, fromAccount.ID)

		toAccountResult := result.ToAccount
		require.NotEmpty(t, toAccountResult)
		require.Equal(t, toAccountResult.ID, toAccount.ID)

		// Check account balance
		fmt.Println(">> tx:", fromAccountResult.Balance, toAccountResult.Balance)
		diff1 := fromAccount.Balance - fromAccountResult.Balance
		diff2 := toAccountResult.Balance - toAccount.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	// Check the final updated balance
	updatedAccount1, err := store.GetAccount(context.Background(), fromAccount.ID)
	require.NoError(t, err)

	updatedAccount2, err := store.GetAccount(context.Background(), toAccount.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)

	require.Equal(t, fromAccount.Balance-int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, toAccount.Balance+int64(n)*amount, updatedAccount2.Balance)
}

func TestTransferTxDeadlock(t *testing.T) {
	store := NewStore(testDB)

	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	fmt.Println(">> before:", fromAccount.Balance, toAccount.Balance)

	// Run concurrent transactions
	n := 10
	amount := int64(100000)

	errs := make(chan error)

	for i := 0; i < n; i++ {
		fromAccountId := fromAccount.ID
		toAccountId := toAccount.ID

		if i%2 == 1 {
			fromAccountId = toAccount.ID
			toAccountId = fromAccount.ID
		}

		go func() {
			ctx := context.Background()
			_, err := store.TransferTx(ctx, TransferTxParams{
				FromAccountID: fromAccountId,
				ToAccountID:   toAccountId,
				Amount:        amount,
			})

			errs <- err
		}()
	}

	// Check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	// Check the final updated balance
	updatedAccount1, err := store.GetAccount(context.Background(), fromAccount.ID)
	require.NoError(t, err)

	updatedAccount2, err := store.GetAccount(context.Background(), toAccount.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)

	require.Equal(t, fromAccount.Balance, updatedAccount1.Balance)
	require.Equal(t, toAccount.Balance, updatedAccount2.Balance)
}
