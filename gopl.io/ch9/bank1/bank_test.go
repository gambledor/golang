// Package bank_test provides test for bank
package bank_test

import (
	"fmt"
	"github.com/gambledor/golang/gopl.io/ch9/bank1"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})
	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()
	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transaction
	<-done
	<-done
	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestBandWithdraw(t *testing.T) {
	done := make(chan struct{})
	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()
	// Bob
	go func() {
		if bank.Withdraw(100) {
			fmt.Println("success")
		} else {
			fmt.Println("Insufficient founds.")
		}
		done <- struct{}{}
	}()

	// Wait for both transaction
	<-done
	<-done
	if got, want := bank.Balance(), 400; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
