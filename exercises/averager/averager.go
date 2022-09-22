package averager

// Variadic functions!
func Averager(numbers ...int) float32 {
	var sum float32 = 0
	for i := 0; i < len(numbers); i++ {
		sum += float32(numbers[i])
	}
	return sum / float32(len(numbers))
}
