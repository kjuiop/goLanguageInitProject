package main

import (
	"sync"
	"time"
)

var mutex sync.Mutex

type Account struct {
	Balance int
}

func main() {
	var wg sync.WaitGroup

	account := &Account{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()
	defer mutex.Unlock()
	account.Balance += 1000
	time.Sleep(time.Millisecond)
}
