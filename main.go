package main

import (
	"fmt"
	"sync"
)

type User struct {
	Name   string
	Gender string
}

func (u *User) change() {
	if u.Gender == "F" {
		u.Gender = "Female"
	} else if u.Gender == "M" {
		u.Gender = "Male"
	}
}

var users = []User{
	{Name: "John", Gender: "F"},
	{Name: "Jack", Gender: "F"},
	{Name: "May", Gender: "M"},
	{Name: "Sonia", Gender: "M"},
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < len(users); i++ {
		wg.Add(1)
		go func(user *User) {
			user.change()

			wg.Done()
		}(&users[i])
	}

	wg.Wait()

	fmt.Println(users)
}
