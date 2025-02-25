package maps_go

import "fmt"

func Maps() {
	ages := make(map[string]int)
	ages["Dhruva"] = 20
	ages["Shivang"] = 20
	ages["Amit"] = 20

	delete(ages, "Dhruva")
	delete(ages, "Dhruva")

	_, exists := ages["Sachin"]
	fmt.Println(exists)

	fmt.Println(len(ages))

	fmt.Println(ages)
}

type Key struct {
	Path, Country string
}

func StructAsKeys() {
	hits := map[Key]int{}
	hits[Key{"/", "vn"}]++
	hits[Key{"/", "nice"}]++
	fmt.Println(hits)
}
