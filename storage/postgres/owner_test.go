package postgres_test

import (
	"github.com/MuhammadyusufAdhamov/booking/storage/repo"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
	"testing"
)

func createOwner(t *testing.T) *repo.Owner {
	owner, err := strg.Owner().Create(&repo.Owner{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Email:     faker.Email(),
		Password:  faker.Password(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, owner)

	return owner
}

func TestGetOwner(t *testing.T) {
	c := createOwner(t)

	owner, err := strg.Owner().Get(c.ID)
	require.NoError(t, err)
	require.NotEmpty(t, owner)
}

func TestCreateOwner(t *testing.T) {
	createOwner(t)
}

func TestGetAllOwners(t *testing.T) {
	createOwner(t)

	result, err := strg.Owner().GetAll(&repo.GetAllOwnersParams{
		Limit: 3,
		Page:  1,
	})
	require.NoError(t, err)
	require.GreaterOrEqual(t, int(result.Count), 1)
}

func TestUpdateOwner(t *testing.T) {
	c := createOwner(t)

	c.FirstName = faker.FirstName()

	owner, err := strg.Owner().Update(c)
	require.NoError(t, err)
	require.NotEmpty(t, owner)
	require.Equal(t, owner.FirstName, c.FirstName)
}

func TestDeleteOwner(t *testing.T) {
	c := createOwner(t)

	err := strg.Owner().Delete(c.ID)
	require.NoError(t, err)
}
