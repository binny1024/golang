package main

func main() {
	a, b := 1, 2
	println("a = ", a, "b = ", b)
	swapPtr(&a, &b)
	println("a = ", a, "b = ", b)
}

/**
形参为指针，该函数接收实参的地址作为参数
*/
func swapPtr(a, b *int) {
	*a, *b = *b, *a
}
