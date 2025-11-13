# Bank Account Kata

A hands-on coding exercise (kata) to practice Go fundamentals including structs, methods, error handling, and concurrency primitives.

## Overview

This kata implements a thread-safe bank account with deposit and withdraw functionality. It demonstrates common Go patterns and includes comprehensive tests for both correctness and concurrency safety.

## Features

- ✅ Account creation with zero balance
- ✅ Deposit and withdrawal operations
- ✅ Input validation (no negative or zero amounts)
- ✅ Insufficient funds checking
- ✅ Thread-safety with `sync.Mutex`
- ✅ Comprehensive test coverage including race condition tests

## Requirements

Create a `BankAccount` type with the following behavior:

1. **NewAccount()** - Creates a new account with a balance of 0
2. **Deposit(amount int)** - Adds money to the account
   - Returns an error if amount is non-positive
3. **Withdraw(amount int)** - Removes money from the account
   - Returns an error if amount is non-positive
   - Returns an error if there are insufficient funds
4. **Balance()** - Returns the current balance

## Go Concepts Practiced

### Core Concepts
- Structs and methods
- Pointer receivers
- Error handling with `error` interface
- Writing tests with the `testing` package

### Concurrency Concepts
- Race conditions and data races
- `sync.Mutex` for mutual exclusion
- `sync.WaitGroup` for goroutine synchronization
- Go's race detector (`-race` flag)

## Running Tests

Run all tests:
```bash
go test ./bankaccount
```

Run with verbose output:
```bash
go test -v ./bankaccount
```

Run with race detector (detects concurrent access issues):
```bash
go test -race ./bankaccount
```

Run only concurrency tests:
```bash
go test -v ./bankaccount -run Concurrent
```

## Implementation Tips

### Basic Implementation
- Use a struct to hold the account state
- Methods with pointer receivers: `func (a *BankAccount) MethodName()`
- Create errors using `fmt.Errorf("message")` for formatted error messages
- Return `nil` for error when there's no error

### Thread-Safety
- Add a `sync.Mutex` field to the struct
- Lock before accessing shared state: `a.mu.Lock()`
- Always unlock when done: `defer a.mu.Unlock()`
- Lock in ALL methods that access the balance (including `Balance()`)

## Test Suite

The kata includes several types of tests:

- **Basic functionality** - Account creation, balance checking
- **Business logic** - Deposits, withdrawals, validation
- **Error cases** - Negative amounts, insufficient funds
- **Concurrency** - Race condition detection with multiple goroutines

The concurrency tests use 1000 goroutines performing 100 operations each to ensure thread-safety under high contention.

## Learning Goals

1. **Understand structs and methods** - Build a simple but complete data structure
2. **Master error handling** - Return and validate errors appropriately
3. **Write testable code** - Create comprehensive test coverage
4. **Recognize race conditions** - Understand when concurrent access is unsafe
5. **Implement thread-safety** - Use mutexes to protect shared state
6. **Use Go's tooling** - Leverage the race detector to find concurrency bugs

## Challenge

Try implementing the kata in stages:

1. **Stage 1**: Basic implementation without concurrency (tests will fail with `-race`)
2. **Stage 2**: Add `sync.Mutex` to make it thread-safe (all tests pass with `-race`)

This progression helps understand why synchronization is necessary and how to implement it correctly.
