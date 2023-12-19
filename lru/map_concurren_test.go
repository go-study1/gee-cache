package lru

import (
	"fmt"
	"sync"

	"testing"
)

var set = make(map[int]bool, 0)
var m sync.Mutex

func printOnce(num int, callback func()) {
	defer callback()
	m.Lock()
	if _, exist := set[num]; !exist {
		fmt.Println(num)
	}
	set[num] = true
	m.Unlock()
}

func TestConcurren(t *testing.T) {
	var group sync.WaitGroup
	group.Add(10)
	for i := 0; i < 10; i++ {
		go printOnce(100, func() {
			group.Done()
		})
	}
	group.Wait()
}

type Counter struct {
	count int
}
type PrintCounter interface {
	PrintCounter1()
	PrintCounter2()
}

func (c Counter) PrintCounter1() {
	c.count++
	fmt.Printf("count=%d, p:%p\r\n", c.count, &c)
}
func (c *Counter) PrintCounter2() {
	c.count++
	fmt.Printf("count=%d, p:%p\r\n", c.count, &c)
}

func TestCounter(t *testing.T) {
	var c PrintCounter = &Counter{}
	c.PrintCounter1()
	c.PrintCounter1()
	c.PrintCounter2()
	c.PrintCounter2()

}
