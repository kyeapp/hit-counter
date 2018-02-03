package main

import (
	"testing"
	"sync"
)

func run_hit_test(t *testing.T, c *Counter, max int) {
	var wg sync.WaitGroup
    wg.Add(max)

	for i := 0; i < max; i++ {
		go func() {
			defer wg.Done()
			c.Add_count()
		}()
	}
	wg.Wait()

	final_count := c.Count()

	if final_count != max {
       t.Errorf("X Hit count was incorrect, got: %d, want: %d.", final_count, max)
    } else {
    	t.Logf("O Hit count was correct, got: %d, want: %d.", final_count, max)
    }

	c.Reset_count()
}

func TestHit(t *testing.T) { 
	testCounter := new(Counter)

	test_hit := []int{10, 100, 1000, 10000, 100000, 1000000, 10000000}

	for _, i := range(test_hit) {
		run_hit_test(t, testCounter, i)
	}

}