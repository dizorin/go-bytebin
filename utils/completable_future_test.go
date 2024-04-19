package utils

import (
	"testing"
	"time"
)

func TestNewCompletableFuture(t *testing.T) {
	cf := NewCompletableFuture[int]()

	go func() {
		time.Sleep(time.Second)
		cf.Complete(10, nil)
	}()

	if cap(cf.signal) != 1 {
		t.Errorf("Expected signal channel size to be 1, but got %d", cap(cf.signal))
	}

	_, _ = cf.CompleteGet()

	value, err := cf.CompleteGet()
	if err != nil {
		t.Fatalf("Expected error: %v", err)
	}

	if value != 10 {
		t.Fatalf("Expected value to be 10, but got %d", value)
	}
}

func TestNewCompletableFutureWithTimeout(t *testing.T) {
	cf := NewCompletableFuture[int]()

	go func() {
		time.Sleep(time.Second)
		cf.Complete(20, nil)
		cf.Complete(10, nil)
	}()

	value, err := cf.CompleteGetWithTimeout(time.Second * 2)
	if err != nil {
		t.Fatalf("Expected error: %v", err)
	}

	if value != 20 {
		t.Fatalf("Expected value to be 20, but got %d", value)
	}
}
