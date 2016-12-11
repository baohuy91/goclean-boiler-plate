package jwtauth

// This salt is used as a salt combination for hashing password
var sysSalt string

func SetSysSalt(authSalt string) {
	sysSalt = authSalt
}

func GetSysSalt() string {
	return sysSalt
}
