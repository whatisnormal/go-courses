//Holds utilitary functions to handle dogs.
package dog

// Converts human years to dog years.
func Years(humanYears int) int {
	return humanYears * 7
}

// YearsTwo converts human years to dog years.
func YearsTwo(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count += 7
	}
	return count
}
