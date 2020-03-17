package benchmark

func slice1() []int {
	m := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		m = append(m, i)
	}
	return m
}
func slice2() []int {
	m := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}
	return m
}
func slice3() []int {
	m := [1000]int{}
	for i := 0; i < 1000; i++ {
		m[i] = i
	}
	return m[:]
}
