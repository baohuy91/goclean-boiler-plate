package domain

type Reason int

const (
	REPO_INTERNAL_ERR Reason = iota
	REPO_CONSTRAINT_ERR
)

type Error struct {
	OriginalErr error
	Reason      Reason
}

func (e Error) Error() string {
	switch e.Reason {
	case REPO_INTERNAL_ERR:
		return "Internal execution in database gateway"
	case REPO_CONSTRAINT_ERR:
		return "Constraint violation for database"
	default:
		return "Unknown Error"
	}
}

func NewRepoInternalErr(err error) error {
	return Error{
		OriginalErr: err,
		Reason:      REPO_INTERNAL_ERR,
	}
}
