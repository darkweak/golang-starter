package demo_tests

func divide2By_Solution(myNumber int) int {
	if myNumber < 0 {
		return 0
	}
	if myNumber == 0 {
		return 9999999
	}

	return 2 / myNumber
}
