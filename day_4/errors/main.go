package learn_errors

import "fmt"

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	var cost int = 0
	costToCustomer, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0, err
	}

	cost += costToCustomer

	costToSpouse, err2 := sendSMS(msgToSpouse)
	if err2 != nil {
		return 0, err2
	}

	cost += costToSpouse
	return cost, nil
}

func sendSMS(message string) (int, error) {
	const maxSMSLength = 25
	const costPerChar = 2
	if len(message) > maxSMSLength {
		return 0, fmt.Errorf("Message is too long")
	}
	return len(message) * costPerChar, nil
}

func Errors() {
	cost, err := sendSMSToCouple("Hello", "Hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Total cost is: ", cost)
	}

	cost, err = sendSMSToCouple("Hello, this is a very long message", "Hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Total cost is: ", cost)
	}
}
