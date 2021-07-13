package repo

// UserRepository represent repository of users
type UserRepository interface {
	FindAll() ([]User, error)
}

// User user entry
type User struct {
	Id   int
	Name string
	Age  int
}
