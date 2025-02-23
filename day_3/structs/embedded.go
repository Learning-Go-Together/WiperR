package structs

import "fmt"

type car_e struct {
	brand string
	model string
}

type truck struct {
	car_e
	bedSize int
}

func EmbeddedStructs() {
	truck := truck{
		car_e: car_e{
			brand: "Toyota",
			model: "Camry",
		},
		bedSize: 20,
	}

	fmt.Println(truck.model)
	fmt.Println(truck.car_e.model)
	fmt.Println(truck.bedSize)
}
