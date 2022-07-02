package batch

import (
	"fmt"
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
	Userchan := make(chan []user, pool)
	//Userchan := make(chan struct{}, pool)
	var wg sync.WaitGroup
	var i int64
	for i = 0; i < n; i++ {
		wg.Add(1)
		go func(nextId int64) {
			user := getOne(nextId)
			res = append(res, user)
			Userchan <- res
			fmt.Println(res)
			wg.Done()
			//close(Userchan)
		}(i)
		wg.Wait()
	}
	return nil
}
