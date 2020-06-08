package service

// User test struct
type User struct {
	Name string
	Age  int
}

// GetName test
func (user *User) GetName() string {
	return user.Name
}

// GetAge test
func (user *User) GetAge() int {
	return user.Age
}
