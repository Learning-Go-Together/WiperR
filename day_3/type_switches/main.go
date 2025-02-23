package type_switches

import "fmt"

func Switch(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("%d\n", v)
	case float64:
		fmt.Printf("%.2f\n", v)
	default:
		fmt.Printf("%v\n", v)
	}
}

func TypeSwitches() {
	Switch(1)
	Switch(2.0)
	Switch(struct{ v string }{v: "Shivang"})

}
