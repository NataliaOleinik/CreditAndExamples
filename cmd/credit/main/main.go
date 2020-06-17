package main

import (
	"fmt"
	_ "fmt"
	"github.com/NataliaOleinik/CreditAndExamples/pkg/credit"
)

func main() {
	amount := int64(100_00_00_00)
	period := int64(36)
	percent := int64(20)
	monthlyPayment, overPayment, totalAmount := credit.Calculate(amount, period, percent)
	fmt.Println(monthlyPayment, overPayment, totalAmount)
}
