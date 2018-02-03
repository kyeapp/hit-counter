/*
 * Now the hit counter is atomic, so the counter is keeping track of hits accurately now. 
 * The issue here is that as the number throughput increase the longer it takes to complete the test.
 * scalability is the issue here.
 */

package main

import (
	"sync/atomic"
	"time"
)

type Counter struct {
	hitCount uint64
}

func (c *Counter) Add_count() {
	atomic.AddUint64(&c.hitCount, 1)
}

func (c *Counter) Count() uint64 {
	time.Sleep(time.Second)
	return c.hitCount
}

func (c *Counter) Reset_count() {
	c.hitCount = 0
}

func main() {

}
