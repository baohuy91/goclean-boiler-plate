package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {

}

func TestUserRepoImpl_Get_WithUserExist_ExpectData(t *testing.T) {
	mockDbGateway := &MockDbGateway{
		ModifiedParam1: &UserModel{
			Id: "12",
		},
		Result1: nil,
	}
	sut := userRepoImpl{dbGateway: mockDbGateway}

	user, err := sut.Get("")

	assert.Nil(t, err)
	assert.Equal(t, "12", user.Id)
}
