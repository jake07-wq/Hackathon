package piscine

func Unmatch(a []int) int {
	for _, target := range a {
		count := 0
		for _, value := range a {
			if value == target {
				count++
			}
		}
		if count%2 != 0 {
			return target
		}
	}
	return -1
}
