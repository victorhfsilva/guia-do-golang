package model

type Account interface {
	Withdraw(value float64) error
	Deposit(value float64) error
	Transfer(value float64, destiny Account) error
}
