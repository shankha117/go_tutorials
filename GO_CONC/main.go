package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	cacheCh := make(chan Book)
	dbCh := make(chan Book)
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			if b, ok := queryCache(id, m); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, m, cacheCh)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			if b, ok := queryDatabase(id, m); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, m, dbCh)
		// create one go routine per query to handel the response
		go func(cacheCh, dbCh <-chan Book) {
			select {
			case b := <-cacheCh:
				fmt.Println("From Cache")
				fmt.Println(b)
				<-dbCh
			case b := <-dbCh:
				fmt.Println("From DataBase")
				fmt.Println(b)

			}

		}(cacheCh, dbCh)

		time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
}

func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	m.RLock()
	b, ok := cache[id]
	// fmt.Println("DATA CALLED IN CACHE FOR ID", id, ok)
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (Book, bool) {
	time.Sleep(300 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			m.Lock()
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}
	return Book{}, false
}
