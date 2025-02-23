package billing_system

import "fmt"

// The number of messages sent
// The cost per message
// Thresholds for discounts

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	baseBill := calculateBaseBill(costPerMessage, numMessages)
	discountRate := calculateDiscountRate(numMessages)
	return baseBill * (1 - discountRate)
}

func calculateDiscountRate(messagesSent int) float64 {
	if messagesSent > 5000 {
		return 0.2
	} else if messagesSent >= 1000 {
		return 0.1
	} else {
		return 0
	}
}

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}

func BillingSystem() {
	costPerMessage := 0.02
	numMessages := 10000

	finalBill := calculateFinalBill(costPerMessage, numMessages)
	fmt.Println("Final Bill: ", finalBill)
}
