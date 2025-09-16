package code

// helper of find gcd
func gcd(a int, b int) int {
	for b != 0 {
		var temp int = b
		b = a % b
		a = temp
	}

	return  a
}

// helper to find lcm
func lcm(a int, b int) int {
	return (a * b) / gcd(a, b)
}

// main function
func replaceNonCoprimes(nums []int) []int {
	stack := []int{}

	for _,num := range nums {
		for len(stack) > 0 {
			top := stack[len(stack) - 1]
			g := gcd(top, num)

			if g > 1 {
				num = lcm(top, num)
				stack = stack[:len(stack) - 1]
			} else {
				break
			}
		}

		stack = append(stack, num)
	}

	return  stack
}