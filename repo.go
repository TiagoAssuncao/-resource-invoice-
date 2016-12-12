package main

import "fmt"

var currentId int

var users Users
var invoices Invoices

// Give us some seed data
func init() {
    RepoCreateUser(User{FirstName: "Tiago"})
    RepoCreateUser(User{FirstName: "Hervylla"})
    RepoCreateInvoice(Invoice{Document: "ASfdsafHervylla"})
    RepoCreateInvoice(Invoice{Document: "ASfdsafHervylla"})
    RepoCreateInvoice(Invoice{Document: "ASfdsafHervylla"})
    RepoCreateInvoice(Invoice{Document: "ASfdsafHervylla"})
}

func RepoFindUser(id int) User {
    for _, u := range users {
        if u.Id == id {
            return u
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
    return u
}

func RepoCreateInvoice(i Invoice) Invoice{
    currentId += 1
    i.Id = currentId
    invoices = append(invoices , i)
    return i
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
