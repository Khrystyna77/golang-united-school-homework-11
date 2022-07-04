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
			var mx sync.Mutex
			mx.Lock()
			fmt.Println(user1.ID)
			res = append(res, user1)
			mx.Unlock()
			Userchan <- res
			mes := <-Userchan
			//var res2 []user
			//res2 = append(res2, user1[i])
			fmt.Println(mes)
			wg.Done()
			//close(Userchan)
		}(i)

		wg.Wait()

	}
	return nil
}

// func main() {
// 	getOne(1)
// 	fmt.Println(getBatch(1, 1))

// }
