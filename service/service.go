package service

// User test struct
type User struct {
	name string
	age  int
}

// GetName test
func (user *User) GetName() string {
	return user.name
}

// GetAge test
func (user *User) GetAge() int {
	return user.age
}
