package repository

type AuthRepo interface {
	Get(userId string) (*Auth, error)
	CreateAuthByEmailAndHashPass(uid, email, hashPash, salt string) (string, error)
}

func NewAuthRepo() AuthRepo {
	return &authRepoImpl{}
}

type authRepoImpl struct {
}

func (r *authRepoImpl) Get(userId string) (*Auth, error) {
	// TODO: implement here
	return &Auth{}, nil
}

func (r *authRepoImpl) CreateAuthByEmailAndHashPass(uid, email, hashPash, salt string) (string, error) {
	// TODO: implement here
	return "", nil
}
