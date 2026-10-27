package main

var users []string

func AddUser(user string) {
	users = append(users, user)
}

func main() {
	AddUser("Alice")
	AddUser("Bob")
}
