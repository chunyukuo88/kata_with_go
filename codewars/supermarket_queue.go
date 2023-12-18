package kata
// https://codewars.com/kata/57b06f90e298a7b53d000a86/train/go

func calcSingleRegisterTime(customers []int) int {
	var result int
	for _, val := range customers {
	  result += val
	}
	return result
  }
  
  func calcMultipleRegisterTime(customers []int, n int) int {
	var tills []int
	for i := 0; i < n; i++ {
	  tills = append(tills, 0)
	}
	customersClone := []int{customers...}
	for i, cust := range customersClone {
		
	}
	
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
  