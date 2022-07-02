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
		go func(userId int64) {
			user1 := getOne(userId)
			res = append(res, user1)
			Userchan <- res
			mes := <-Userchan
			fmt.Println(mes)
			wg.Done()
			//close(Userchan)
		}(i)

		wg.Wait()
	}
	return nil
}
