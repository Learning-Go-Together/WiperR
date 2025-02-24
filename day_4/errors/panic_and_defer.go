package learn_errors

func PanicAndDefer() {
	defer func() {
		if r := recover(); r != nil {
			println("Recovered from panic:", r)
		}
	}()

	panic("Panic!")
	println("This will not be executed")
}
