package calculator

func Addition(a, b int) int {
	return a + b
}

func Substraction(a, b int) int {
	return a - b
}

func Multiplication(a, b int) int {
	return a * b
}

func Division(a, b int) interface{} {
	if b == 0 {
		return "can't divide by 0"
	}
	return a / b
}
