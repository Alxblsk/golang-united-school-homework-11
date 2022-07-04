package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	semaphore := make(chan struct{}, pool)

	var i int64 = 0
	for i = 0; i < n; i++ {
		wg.Add(1)

		semaphore <- struct{}{}

		go func(id int64) {
			defer wg.Done()

			value := getOne(id)
			<-semaphore

			mu.Lock()
			res = append(res, value)
			mu.Unlock()
		}(i)

	}

	wg.Wait()

	return res
}
