package singleflight

import "sync"

type call struct {
	wg    sync.WaitGroup
	value interface{}
	err   error
}

type Group struct {
	mu       sync.Mutex
	callings map[string]*call
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.callings == nil {
		g.callings = map[string]*call{}
	}
	if call, ok := g.callings[key]; ok {
		call.wg.Wait()
		return call.value, call.err
	}
	c := new(call)
	c.wg.Add(1)
	g.callings[key] = c
	c.value, c.err = fn()
	c.wg.Done()
	delete(g.callings, key)
	return c.value, c.err
}
