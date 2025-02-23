package deferfn

func DeferTest() {
	defer println("This is deffered")
	println("Hello World")
}
