package main

import "sync"

func updateMessage(s string, msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg = s
}

func updateMessageMutex(override string, msg *string, m *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Lock()
	*msg = override
	m.Unlock()
}
