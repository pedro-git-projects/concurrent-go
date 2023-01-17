package main

import (
	"sync"
	"testing"
)

func TestUpdateMessageMutex(t *testing.T) {
	override := "Goodbye, cruel world!"
	msg := "Hello, world"
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go updateMessageMutex(override, &msg, &mutex, &wg)
	wg.Wait()

	if msg != override {
		t.Errorf("Expected %s but got %s", override, msg)
	}
}
