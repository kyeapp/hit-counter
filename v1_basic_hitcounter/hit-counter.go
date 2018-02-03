/* 
 * With a setup like this it doesn't account for the race conditions of the hit counter 
 * being accessed too fast.
 */

package main

import (
//"fmt"
)

type Counter struct {
	hitCount int
}

func (c *Counter) Add_count() {
	c.hitCount++
}

func (c *Counter) Count() int {
	return c.hitCount
}

func (c *Counter) Reset_count() {
	c.hitCount = 0
}

func main() {

}
