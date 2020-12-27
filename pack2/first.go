package pack2

// function, which takes a string as
// argument and return the reverse of string.
func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
