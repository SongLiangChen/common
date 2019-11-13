package common

import (
	"sync"
)

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		cb()
		w.Done()
	}()
}

func (w *WaitGroupWrapper) WrapWithParams(cb func(...interface{}), p ...interface{}) {
	w.Add(1)
	go func() {
		cb(p...)
		w.Done()
	}()
}
