package type_assertion

import "fmt"

func TypeAssertion() {
	var value interface{} = "Shivang Rathore"
	// Safe guard
	s, ok := value.(int)
	if ok {
		fmt.Println(s)
	}

	// Will panic on error
	s1 := value.(string)
	fmt.Println(s1)
}
