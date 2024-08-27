package users

type WrongUsernameOrPasswordError struct{}

func (w *WrongUsernameOrPasswordError) Error() string {
	return "Incorrect username or password"
}
