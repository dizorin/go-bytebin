package utils

import (
	"fmt"
	"sync"
	"time"
)

type CompletableFuture[T any] struct {
	err       error
	value     T
	signal    chan struct{}
	mutexRead sync.Mutex
	onceWrite sync.Once
}

func NewCompletableFuture[T any]() *CompletableFuture[T] {
	return &CompletableFuture[T]{
		signal:    make(chan struct{}, 1),
		onceWrite: sync.Once{},
	}
}

func (cf *CompletableFuture[T]) Complete(v T, err error) {
	cf.onceWrite.Do(func() {
		cf.value = v
		cf.err = err
		cf.signal <- struct{}{}
		close(cf.signal)
	})
}

func (cf *CompletableFuture[T]) CompleteGet() (T, error) {
	cf.await()
	return cf.value, cf.err
}

func (cf *CompletableFuture[T]) CompleteGetWithTimeout(duration time.Duration) (any, error) {
	err := cf.awaitWithTimeout(duration)
	if err != nil {
		return nil, err
	}
	return cf.value, cf.err
}

func (cf *CompletableFuture[T]) await() {
	cf.mutexRead.Lock()
	defer cf.mutexRead.Unlock()
	<-cf.signal
}

func (cf *CompletableFuture[T]) awaitWithTimeout(duration time.Duration) error {
	cf.mutexRead.Lock()
	defer cf.mutexRead.Unlock()

	timer := time.NewTimer(duration)
	select {
	case <-timer.C:
		return fmt.Errorf("timeout")
	case <-cf.signal:
		timer.Stop()
	}
	return nil
}
