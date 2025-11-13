package bankaccount

import (
	"sync"
	"testing"
)

func setupAccount() *BankAccount {
	return NewAccount()
}

func TestNewAccountHasZeroBalance(t *testing.T) {
	account := setupAccount()

	if account == nil {
		t.Errorf("expected account to be defined")
	}

	if account.Balance() != 0 {
		t.Errorf("expected account to have balance of 0")
	}
}

func TestWithdrawFromEmptyAccountIsNotAllowed(t *testing.T) {
	account := setupAccount()

	err := account.Withdraw(5)

	if err == nil {
		t.Errorf("expected withdraw to return error")
	}
}

func TestDepositWillIncreaseBalanceByCorrectAmount(t *testing.T) {
	account := setupAccount()

	initialBalance := account.Balance()

	err := account.Deposit(100)

	balanceAfterDeposit := account.Balance()

	if err != nil {
		t.Errorf("expected successful deposit")
	}

	if initialBalance+100 != balanceAfterDeposit {
		t.Errorf("expected balance to be: %d instead got: %d", initialBalance+100, balanceAfterDeposit)
	}
}

func TestWithdrawWillReduceBalanceByRequestedAmount(t *testing.T) {
	account := setupAccount()

	account.Deposit(100)
	withdrawError := account.Withdraw(75)

	balanceAfterOperation := account.Balance()

	if withdrawError != nil {
		t.Errorf("expected withdraw successfuly without errors")
	}

	if balanceAfterOperation != 25 {
		t.Errorf("expected balance to equal 25 but instead got: %v", balanceAfterOperation)
	}
}

func TestWithdrawAndDepositNegativeAmountIsNotAllowed(t *testing.T) {
	account := setupAccount()

	balanceBeforeOperation := account.Balance()

	withdrawError := account.Withdraw(-100)
	depositError := account.Deposit(-20)

	balanceAfterOperation := account.Balance()

	if depositError == nil {
		t.Errorf("expected deposit to return error")
	}

	if withdrawError == nil {
		t.Errorf("expected withdraw to return error")
	}

	if balanceAfterOperation != balanceBeforeOperation {
		t.Errorf("expected balance to remain the same")
	}
}

func TestConcurrentDepositsAndWithdrawals(t *testing.T) {
	// Run the test multiple times to increase chance of catching the race
	for attempt := 0; attempt < 10; attempt++ {
		account := setupAccount()
		startingBalance := 50000
		account.Deposit(startingBalance)

		var wg sync.WaitGroup
		numGoroutines := 1000
		operationsPerGoroutine := 100

		// Half deposit, half withdraw
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			if i%2 == 0 {
				// Depositors
				go func() {
					defer wg.Done()
					for j := 0; j < operationsPerGoroutine; j++ {
						account.Deposit(1)
					}
				}()
			} else {
				// Withdrawers
				go func() {
					defer wg.Done()
					for j := 0; j < operationsPerGoroutine; j++ {
						account.Withdraw(1)
					}
				}()
			}
		}

		wg.Wait()

		// Net change: 500 goroutines deposit, 500 withdraw, 100 ops each = 0 net change
		expectedBalance := startingBalance
		actualBalance := account.Balance()

		if actualBalance != expectedBalance {
			t.Fatalf("race condition detected on attempt %d: expected balance %d, got %d", attempt+1, expectedBalance, actualBalance)
		}
	}
}
