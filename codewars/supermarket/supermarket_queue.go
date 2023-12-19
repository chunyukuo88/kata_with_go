package supermarket

// https://codewars.com/kata/57b06f90e298a7b53d000a86/train/go

// This is a bloody mess. Come back to this later... feels like Golang isn't made for these kinds of kata.

func calcSingleRegisterTime(customers []int) int {
	var result int
	for _, val := range customers {
		result += val
	}
	return result
}

func createTills(n int) []int {
	var tills []int
	for i := 0; i < n; i++ {
		tills = append(tills, 0)
	}
	return tills
}

func calcMultipleRegisterTime(customers []int, n int) int {
	var result int
	tills := createTills(n)

	var tillsIndex int
	for i, customerInterval := range customers {
		if i == len(customers)-1 {
			tillsIndex = 0
		}
		if i < (len(customers) - 1) {
			tills[tillsIndex+i] += customerInterval
		} else {
			tills[tillsIndex+i] += customerInterval
		}
	}
	for _, totalTime := range tills {
		result += totalTime
	}
	return result
}

func QueueTime(customers []int, n int) int {
	var result int
	if len(customers) == 0 {
		return 0
	} else if n == 1 {
		result = calcSingleRegisterTime(customers)
	} else {
		result = calcMultipleRegisterTime(customers, n)
	}
	return result
}
