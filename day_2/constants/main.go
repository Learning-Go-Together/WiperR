package constants

import "fmt"

func Constants() {
	const premiumPlan = "Premium"
	// premiumPlan = "Basic"
	const basicPlan = "Basic"
	fmt.Println(premiumPlan)
	fmt.Println(basicPlan)
}

func StaticConstants() {
	const firstName = "Shivang"
	const lastName = "Rathore"
	const fullName = firstName + " " + lastName
	fmt.Println(fullName)
}
