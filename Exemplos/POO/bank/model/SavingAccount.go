package model

import "errors"

type SavingsAccount struct {
	holder   Holder
	agency   int
	number   int
	value    float64
	interest float64
}

func CreateSavingsAccount(holder Holder, agency int, number int, value float64) SavingsAccount {
	return SavingsAccount{
		holder: holder,
		agency: agency,
		number: number,
		value:  value,
	}
}

func (a SavingsAccount) GetHolder() Holder {
	return a.holder
}

func (a SavingsAccount) GetAgency() int {
	return a.agency
}

func (a SavingsAccount) GetNumber() int {
	return a.number
}

func (a SavingsAccount) GetValue() float64 {
	return a.value
}

func (a SavingsAccount) GetInterest() float64 {
	return a.interest
}

func (a *SavingsAccount) Withdraw(value float64) error {
	if value > 0 && value <= a.value {
		a.value -= value
	} else {
		return errors.New("withdraw error: invalid value")
	}
	return nil
}

func (a *SavingsAccount) Deposit(value float64) error {
	if value > 0 {
		a.value += value
	} else {
		return errors.New("deposit error: invalid value")
	}
	return nil
}

func (a *SavingsAccount) Transfer(value float64, destiny Account) error {
	if value > 0 && value <= a.value {
		a.value -= value
		destiny.Deposit(value)
	} else {
		return errors.New("transfer error: invalid value")
	}
	return nil
}
