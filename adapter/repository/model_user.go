package repository

type UserModel struct {
	Id    string `gorethink:"id"`
	Name  string `gorethink:"name"`
	Email string `gorethink:"email"`
	Pass  string `gorethink:"pass"`
	Salt  string `gorethink:"salt"`
	CommonModelImpl
}
