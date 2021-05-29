package entity

type User struct {
	id    int
	name  string
	email string
}

func NewUser(id int, name string, email string) *User {
	user := new(User)
	user.id = id
	user.name = name
	user.email = email
	return user
}
