package syncx

import "sync"

type WaitGroup struct {
	sync.WaitGroup
}

func (wg *WaitGroup) Go(f func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		f()
	}()
}
