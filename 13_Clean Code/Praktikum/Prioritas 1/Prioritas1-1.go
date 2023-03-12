package main

type user struct {
	id       int
	username int
	password int
}

// Penamaan menggunakan "t" tidak disarankan, lebih baik "users", karena berisikan beberapa "user"
type userservice struct {
	users []user
}

// Penamaan fungsi sebaiknya kapital pada kata kedua
func (u userservice) getAllusers() []user {
	return u.users
}

// Looping dengan "r" tidak disarankan, lebih baik menggunakan "user" / "usr", karena masih erat dengan "u.users"
func (u userservice) getUserbyid(id int) user {
	for _, usr := range u.users {
		if id == usr.id {
			return usr
		}
	}

	return user{}
}
