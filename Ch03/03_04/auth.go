package main

import "fmt"

type Role string

const (
	Viewer    Role = "viewer"
	Developer Role = "developer"
	Admin     Role = "admin"
)

type User struct {
	Login string
	Role  Role
}

func Promote(u User, r Role) {
	u.Role = r
}

func main() {
	u := User{"elliot", Viewer}
	Promote(u, Admin)
	fmt.Printf("%#v\n", u)
}
