package service

type User struct {
	name string
	age  int
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) GetAge() int {
	return user.age
}
