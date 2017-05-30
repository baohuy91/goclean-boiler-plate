package repository

import (
	"github.com/stretchr/testify/assert"
	"goclean/domain"
	"testing"
)

var sut userRepoImpl

func init() {
	sut = userRepoImpl{}
}

func TestUserRepoImpl_Get_WithUserExist_ExpectData(t *testing.T) {
	sut.dbGateway = MockDbGateway{
		ModifiedParam1: &UserModel{
			Id: "12",
		},
	}

	user, err := sut.Get("")

	assert.Nil(t, err)
	assert.Equal(t, "12", user.Id)
}

func TestUserRepoImpl_GetByEmail_WithUserExist_ExpectData(t *testing.T) {
	sut.dbGateway = MockDbGateway{
		ModifiedParam1: &[]*UserModel{{Id: "45"}},
	}

	user, err := sut.GetByEmail("")

	assert.Nil(t, err)
	assert.Equal(t, "45", user.Id)
}

func TestUserRepoImpl_GetByEmail_WithNoUserNoError_ExpectNil(t *testing.T) {
	sut.dbGateway = MockDbGateway{ModifiedParam1: &[]*UserModel{}}

	user, err := sut.GetByEmail("")

	assert.Nil(t, user)
	assert.Nil(t, err)
}

func TestUserRepoImpl_Create_WithCreatedId_ExpectId(t *testing.T) {
	sut.dbGateway = MockDbGateway{Result1: "12"}

	id, err := sut.Create(domain.User{})

	assert.Equal(t, "12", id)
	assert.Nil(t, err)
}
