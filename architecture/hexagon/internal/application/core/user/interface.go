package user

type IUser interface {
	NewUser(email, password string) (UserData, error)
	ComparePassword(hashedPassword, password string) error
}
