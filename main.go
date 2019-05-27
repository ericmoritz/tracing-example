package main

import (
	"sync"
	"tracing-example/frontend"
	"tracing-example/backend"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go backend.Main(&wg)
	wg.Add(1)
	go frontend.Main(&wg)
	wg.Wait()
}
