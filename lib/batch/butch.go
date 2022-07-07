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
	Userchan := make(chan struct{}, pool)
	//res = make([]user, n)
	//Userchan := make(chan struct{}, pool)
	var mx sync.Mutex
	var i int64
	var wg sync.WaitGroup
	////FOR
	for i = 0; i < n; i++ {
		wg.Add(1)
		Userchan <- struct{}{}
		go func(userId int64) {

			user1 := getOne(userId)

			mx.Lock()
			res = append(res, user1)
			mx.Unlock()
			//Userchan <- res
			//res[i] = getOne(i)
			//<-Userchan
			//var res2 []user
			//fmt.Println(res)
			_, ok := <-Userchan
			if !ok {
				return
				//	fmt.Println("channel close")
			}

			//res2 = append(res2, user1[i])

			wg.Done()

		}(i)
	}
	wg.Wait()
	close(Userchan)
	return
}

// func main() {

// 	fmt.Println(getBatch(10, 1))

// }
