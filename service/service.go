package service

type User struct {
	ID   int
	Name string
	Age  int
}

type Users interface {
	GetUser() *User
}

type UserService struct {
	Users
}

func (us *UserService) GetUser() *User {
	return us.Users.GetUser()
}
