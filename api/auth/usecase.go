package auth

const CtxUserKey = "user"

type UseCase interface {
	SignUp(username, password string) error
	SignIn(username, password string) (string, error)
}
