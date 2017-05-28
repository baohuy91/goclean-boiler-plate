package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {

}

func TestUserRepoImpl_Get_WithUserExist_ExpectData(t *testing.T) {
	mockDbGateway := MockDbGateway{
		ModifiedParam1: &UserModel{
			Id: "12",
		},
	}
	sut := userRepoImpl{dbGateway: mockDbGateway}

	user, err := sut.Get("")

	assert.Nil(t, err)
	assert.Equal(t, "12", user.Id)
}

func TestUserRepoImpl_GetByEmail_WithUserExist_ExpectData(t *testing.T) {
	mockDbGateway := MockDbGateway{
		ModifiedParam1: &[]*UserModel{{Id: "45"}},
	}

	sut := userRepoImpl{dbGateway: mockDbGateway}

	user, err := sut.GetByEmail("")

	assert.Nil(t, err)
	assert.Equal(t, "45", user.Id)
}
