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
	res = make([]user, n)
	var i int64
	wg := new(sync.WaitGroup)
	semaphore := make(chan struct{}, pool)

	wg.Add(int(n))
	for i = 0; i < n; i++ {
		semaphore <- struct{}{}
		go func(j int64, w *sync.WaitGroup) {
			res[j] = getOne(int64(j))
			<-semaphore
			w.Done()
		}(i, wg)
	}
	wg.Wait()
	return res
}
