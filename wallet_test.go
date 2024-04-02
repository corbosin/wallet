package main

import "testing"

func TestDeposit(t *testing.T) {
	var test_wallet Wallet
	tests := []struct {
		input    Bitcoin
		expected Bitcoin
	}{
		{5, 20},
		{60, 75},
		{4, 19},
	}
	for _, test := range tests {
		test_wallet = Wallet{balance: 15}
		test_wallet.Deposit(test.input)
		result := test_wallet.balance
		if result != test.expected {
			t.Errorf("Depositing (%v) makes balance equal to %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestWithdraw(t *testing.T) {
	var test_wallet Wallet
	tests := []struct {
		input    Bitcoin
		expected Bitcoin
	}{
		{5, 10},
		{10, 5},
		{4, 11},
	}
	for _, test := range tests {
		test_wallet = Wallet{balance: 15}
		test_wallet.Withdraw(test.input)
		result := test_wallet.balance
		if result != test.expected {
			t.Errorf("Withdraw (%v) makes balance equal to %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestBalance(t *testing.T) {
	var test_wallet Wallet
	tests := []struct {
		input    Bitcoin
		expected Bitcoin
	}{
		{10, 10},
		{0, 0},
		{15, 15},
	}
	for _, test := range tests {
		test_wallet = Wallet{balance: test.input}
		if test_wallet.balance != test.expected {
			t.Errorf("It is (%v), expected %v", test_wallet.balance, test.expected)
		}
	}
}
