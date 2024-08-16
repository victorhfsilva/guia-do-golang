package main

import (
	"bank/model"
	"fmt"
)

func main() {
	holderCesar := model.CreateHolder("Cesar", "12345678991", "Emperor")
	hodlerAugustus := model.CreateHolder("Augustus", "98765432112", "Emperor")

	accountCesar := model.CreateCheckingAccount(holderCesar, 123, 123456, 902901.10)
	accountAugustus := model.CreateCheckingAccount(hodlerAugustus, 321, 654321, 887054.61)

	transferError := accountCesar.Transfer(150000.0, &accountAugustus)
	if transferError != nil {
		fmt.Println(transferError)
		return
	}

	bankSlip := model.CreateBankSlip(123303.00, &accountCesar)
	paymentError := bankSlip.Pay(&accountAugustus)

	if paymentError != nil {
		fmt.Println(paymentError)
		return
	}

	fmt.Println(accountCesar)
	fmt.Println(accountAugustus)
}
