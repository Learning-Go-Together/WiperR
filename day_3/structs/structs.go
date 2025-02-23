package structs

import "fmt"

type car struct {
	brand      string
	model      string
	doors      int
	milege     int
	frontWheel wheel
	backWheel  wheel
}

type wheel struct {
	radius   int
	material string
}

func Struct() {
	car := car{
		brand:  "Maruti",
		model:  "C17",
		doors:  2,
		milege: 40,
		frontWheel: wheel{
			radius:   5,
			material: "charcoal",
		},
		backWheel: wheel{
			radius:   5,
			material: "charcoal",
		},
	}

	fmt.Println(car)
	car.backWheel.material = "nothing"
	fmt.Println(car.backWheel.material)
}
