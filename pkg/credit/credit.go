package credit

import "math"

func Calculate(amount int64, period int64, percent int64) (monthlyPayment, overPayment, totalAmount int64) {
	var monthlyRate float64 = float64(percent) / 12 / 100
	var coefficient float64 = (monthlyRate) * (math.Pow((1 + monthlyRate), float64(period))) / (math.Pow((1+monthlyRate), float64(period)) - 1)
	monthlyPayment = int64(coefficient * float64(amount))
	totalAmount = monthlyPayment * period
	overPayment = totalAmount - amount
	return monthlyPayment, overPayment, totalAmount
}
