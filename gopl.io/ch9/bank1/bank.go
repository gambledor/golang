// Package bank provides
package bank

type withdrawMsg struct {
	amount int
	out    chan bool // to receive result operation
}

var (
	deposits  = make(chan int)         // send amount to deposit
	balances  = make(chan int)         // receive balance
	withdraws = make(chan withdrawMsg) // get amout from deposit
)

func init() {
	go teller() // start the monitor goroutine
}

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) bool {
	out := make(chan bool)
	msg := withdrawMsg{amount, out}
	withdraws <- msg
	if <-out {
		return true
	}
	return false
}

// Monitor goroutine
func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case withdraw := <-withdraws:
			if balance < withdraw.amount {
				withdraw.out <- false
			} else {
				balance -= withdraw.amount
				withdraw.out <- true
			}
		}
	}
}
