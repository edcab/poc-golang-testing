package exerciseone

/*
Add two numbers
Return the sum of x and y
*/
func Add(x, y int) int {
	return x + y
}

func AddMultiple(numbers ...int) int {

	result := 0

	for _, v := range numbers {
		result += v
	}

	return result

}
