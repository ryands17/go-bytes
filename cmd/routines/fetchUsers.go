package routines

import (
	"fmt"
	"sync"

	"github.com/go-resty/resty/v2"
	"github.com/ryands17/go-bytes/cmd/utils"
)

var client = resty.New()

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func fetchUser(wg *sync.WaitGroup, ch chan<- *User, id int) {
	defer wg.Done()
	var user User

	if _, err := client.R().
		SetResult(&user).
		Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", id)); err != nil {
		ch <- nil
	} else {
		ch <- &user
	}
}

func FetchAllUsers(maxUsers int) {
	var wg sync.WaitGroup
	wg.Add(maxUsers)
	ch := make(chan *User, maxUsers)

	for i := range maxUsers {
		go fetchUser(&wg, ch, i+1)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	users := make([]User, 0, maxUsers)
	for user := range ch {
		if user != nil {
			users = append(users, *user)
		}
	}

	utils.PrintJSON(map[string]any{
		"users": users,
	})
}
