package higherorderfn

func Arthimetic(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

func Add(a, b int) int {
	return a + b
}

func Mul(a, b int) int {
	return a * b
}
