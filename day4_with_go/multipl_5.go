package main

func multiple(n int) int {
	var count int
	if n < 0 {
		return 0
	}
	for i := 1; i <= n; i++ {
		if i%5 == 0 {
			count++
		}

	}
	return count

}
