package model

import "errors"

type BankSlip struct {
	value            float64
	receivingAccount Account
}

func CreateBankSlip(value float64, receivingAccount Account) BankSlip {
	return BankSlip{
		value:            value,
		receivingAccount: receivingAccount,
	}
}

func (b BankSlip) GetValue() float64 {
	return b.value
}

func (b BankSlip) GetReceivingAccount() Account {
	return b.receivingAccount
}

func (b BankSlip) Pay(payerAccount Account) error {
	if payerAccount != b.receivingAccount {
		payerAccount.Transfer(b.value, b.receivingAccount)
	} else {
		return errors.New("payment error: invalid account")
	}
	return nil
}
