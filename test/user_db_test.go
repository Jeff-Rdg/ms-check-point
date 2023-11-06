package test

import (
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"ms-control-point/internal/entity"
	"ms-control-point/internal/infra/database"
	"testing"
)

func TestUserDb_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		t.Error(err)
	}

	user, _ := entity.NewUser("John", "j@j.com", "123456")
	userDb := database.NewUser(db)

	err = userDb.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}

func TestUserDb_FindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		t.Error(err)
	}

	user, _ := entity.NewUser("John", "j@j.com", "123456")
	userDb := database.NewUser(db)

	err = userDb.Create(user)
	assert.Nil(t, err)

	userFound, err := userDb.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}
