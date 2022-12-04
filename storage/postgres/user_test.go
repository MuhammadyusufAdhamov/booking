package postgres_test

import (
	"testing"

	"github.com/MuhammadyusufAdhamov/booking/storage/repo"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func createUser(t *testing.T) *repo.User {
	user, err := strg.User().Create(&repo.User{
		FirstName: faker.FirstName(),
		LastName: faker.LastName(),
		Email: faker.Email(),
		Password: faker.Password(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, user)

	return user
}

func TestGetUser(t *testing.T) {
	c := createUser(t)

	user, err := strg.User().Get(c.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user)
}

func TestCreateUser(t *testing.T) {
	createUser(t)
}

func TestGetAllUsers(t *testing.T) {
	createUser(t)

	result, err := strg.User().GetAll(&repo.GetAllUsersParams{
		Limit: 3,
		Page: 1,
	})
	require.NoError(t, err)
	require.GreaterOrEqual(t, int(result.Count), 1)
}

func TestUpdateUser(t *testing.T) {
	c := createUser(t)

	c.FirstName = faker.FirstName()

	user, err := strg.User().Update(c)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.FirstName, c.FirstName)
}

func TestDeleteUser(t *testing.T) {
	c := createUser(t)

	err := strg.User().Delete(c.ID)
	require.NoError(t, err)
}