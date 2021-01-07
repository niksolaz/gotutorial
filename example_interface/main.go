package main

import (
	"fmt"
)

type Wallet interface {
	MyWallet()
}

type Personal struct {
	Fullname string
	Age      int16
	Address  string
	City     string
}

type Amounted struct {
	Amount float32
	CodID  int32
	Method string
}

func (w Personal) MyWallet() {
	fmt.Printf("Fullname: %s\n", w.Fullname)
	fmt.Printf("Age     : %d\n", w.Age)
	fmt.Printf("Address : %s\n", w.Address)
	fmt.Printf("City    : %s\n", w.City)
}

func (w Amounted) MyWallet() {
	fmt.Printf("Amount  : %.2f\n", w.Amount)
	fmt.Printf("CodID   : %d\n", w.CodID)
	fmt.Printf("Method  : %s\n", w.Method)
}

func main() {
	var personal Personal
	personal.Fullname = "Mario Rossi"
	personal.Age = 44
	personal.Address = "via Tor di Quinto 89"
	personal.City = "Roma"
	personal.MyWallet()

	var amount Amounted
	amount.Amount = 299.59
	amount.CodID = 9807667
	amount.Method = "Bank Wire"
	amount.MyWallet()
}
