package model

import "errors"

type CheckingAccount struct {
	holder Holder
	agency int
	number int
	value  float64
}

func CreateCheckingAccount(holder Holder, agency int, number int, value float64) CheckingAccount {
	return CheckingAccount{
		holder: holder,
		agency: agency,
		number: number,
		value:  value,
	}
}

func (a CheckingAccount) GetHolder() Holder {
	return a.holder
}

func (a CheckingAccount) GetAgency() int {
	return a.agency
}

func (a CheckingAccount) GetNumber() int {
	return a.number
}

func (a CheckingAccount) GetValue() float64 {
	return a.value
}

func (a *CheckingAccount) Withdraw(value float64) error {
	if value > 0 && value <= a.value {
		a.value -= value
	} else {
		return errors.New("withdraw error: invalid value")
	}
	return nil
}

func (a *CheckingAccount) Deposit(value float64) error {
	if value > 0 {
		a.value += value
	} else {
		return errors.New("deposit error: invalid value")
	}
	return nil
}

func (a *CheckingAccount) Transfer(value float64, destiny Account) error {
	if value > 0 && value <= a.value {
		a.value -= value
		destiny.Deposit(value)
	} else {
		return errors.New("transfer error: invalid value")
	}
	return nil
}
