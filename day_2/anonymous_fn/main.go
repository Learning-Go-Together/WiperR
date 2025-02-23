package anonymousfn

func Anonymous() func(a int) int {
	return func(a int) int { return a + 5 }
}
