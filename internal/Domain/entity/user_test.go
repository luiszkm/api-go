package entity_test

import (
	"testing"
	"github.com/luiszkm/api/internal/Domain/entity"
	"github.com/stretchr/testify/assert"
)
func TestNewUSer (t *testing.T){
	user, err := entity.NewUser("Luis", "luis@email.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, user.Name, "Luis")
	assert.Equal(t, user.Email, "luis@email.com" )
}

func TestUserValidatePassword(t *testing.T) {
	user, err :=  entity.NewUser("Luis", "luis@email.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, user.Password , "123456")
}