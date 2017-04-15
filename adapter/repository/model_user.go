package repository

type UserModel struct {
	Id    string
	Name  string
	Email string
	Pass  string
	Salt  string
	CommonModelImpl
}
