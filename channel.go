package benchmark

import (
	"sync"
)

func channel1() []int {
	a := []int{}
	ch := make(chan int, 100)
	go func() {
		for i := 0; i < 100; i++ {
			go func(i int) {
				a = append(a, i)
				ch <- i
			}(i)
		}
	}()
	i := 0
	for range ch {
		i++
		if i >= 100 {
			close(ch)
		}
	}
	return a[:]
}

func sync2() []int {
	a := []int{}
	var sw = new(sync.WaitGroup)
	sw.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			a = append(a, i)
			sw.Done()
		}(i)
	}
	sw.Wait()
	return a[:]
}
