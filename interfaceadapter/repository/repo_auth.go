package repository

type AuthRepo interface {
	Get(userId string) (*Auth, error)
	GetByEmail(email string) (*Auth, error)
	CreateAuthByEmailAndHashPass(uid, email, hashPash, salt string) (string, error)
	// Create or update signed key for user "uid" and at key "aud"
	SaveSignedKey(uid, aud, signedKey string) error
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

func (r *authRepoImpl) GetByEmail(email string) (*Auth, error) {
	// TODO: implement here
	return &Auth{}, nil
}

func (r *authRepoImpl) CreateAuthByEmailAndHashPass(uid, email, hashPash, salt string) (string, error) {
	// TODO: implement here
	return "", nil
}

func (r *authRepoImpl) SaveSignedKey(uid, aud, signedKey string) error {
	return nil
}
