package piscine

func Compact(ptr *[]string) int {
	// Create a temporary slice to hold non-zero values
	var result []string

	// Iterate through the original slice via the pointer
	for _, val := range *ptr {
		if val != "" {
			result = append(result, val)
		}
	}

	// Update the original slice with the compacted version
	*ptr = result

	// Return the count of non-zero elements
	return len(result)
}
