package main

import "fmt"

var currentId int

var users Users

// Give us some seed data
func init() {
	RepoCreateUser(User{Name: "Write presentation"})
	RepoCreateUser(User{Name: "Host meetup"})
}

func RepoFindUser(id int) User {
	for _, u := range users {
		if u.Id == id {
			return t
		}
	}
	// return empty User if not found
	return User{}
}

//this is bad, I don't think it passes race condtions
func RepoCreateUser(u User) User {
	currentId += 1
	u.Id = currentId
	users = append(users, u)
	return t
}

func RepoDestroyUser(id int) error {
	for i, t := range users {
		if t.Id == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find User with id of %d to delete", id)
}
