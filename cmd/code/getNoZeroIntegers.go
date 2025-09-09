package code

func isNonZero(num int) bool {
	for num > 0{
		if num % 10 == 0{
			return  false
		}
		num = num / 10
	}
	return  true
}

func getNoZeroIntegers(n int) []int {
    for i := 1; i < n; i++{
		b := n-1
		if isNonZero(b) && isNonZero(i){
			return []int{i,b}
		}
	}
	return []int{}
}