package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(3)

	start := time.Now()

	type User struct {
		Login string
		Id int64
		Node_id string
		Avatar_url string
		Gravatar_id string
		Url string
		Html_url string
		Followers_url string
		Following_url string
		Gists_url string
		Starred_url string
		Subscriptions_url string
		Organizations_url string
		Repos_url string
		Events_url string
		Received_events_url string
		Type string
		Site_admin bool
		Name string
		Company string
		Blog string
		Location string
		Email string
		Hireable bool
		Bio string
		Public_repos int
		Public_gists int
		Followers int
		Following int
		Created_at string
		Updated_at string
	}

	numComplete := 0

	usernames := []string {"agaro1121", "verasraul", "johanandren", "rveve" , "adamw", "iravid"}

	for _, username := range usernames {
		go func(username string) {
			resp, _ := http.Get("https://api.github.com/users/" + username)
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)

			user := new(User)
			json.Unmarshal(body, &user)

			fmt.Println(body)
			fmt.Println(user)
			numComplete++
		}(username)
	}

	for numComplete < len(usernames) {
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("time taken=", time.Since(start))
}
