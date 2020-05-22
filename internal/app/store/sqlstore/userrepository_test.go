package sqlstore_test

import (
	"github.com/PesTospertnyj/http-rest-api/internal/app/store"
	"github.com/PesTospertnyj/http-rest-api/internal/app/store/sqlstore"
	"testing"

	"github.com/PesTospertnyj/http-rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown("users")
	s := sqlstore.New(db)
	u := model.TestUser()
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown("users")

	s := sqlstore.New(db)
	email := "user@example.org"

	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser()
	u.Email = email

	_ = s.User().Create(u)

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
