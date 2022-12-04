package postgres_test

import (
	"github.com/MuhammadyusufAdhamov/booking/storage/repo"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
	"testing"
)

func createHotel(t *testing.T) *repo.Hotel {
	owner := createOwner(t)

	hotel, err := strg.Hotel().Create(&repo.Hotel{
		HotelName:     faker.NAME,
		HotelLocation: faker.SENTENCE,
		OwnerID:       owner.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, hotel)

	return hotel
}

func TestGetHotel(t *testing.T) {
	c := createHotel(t)

	hotel, err := strg.Hotel().Get(c.ID)
	require.NoError(t, err)
	require.NotEmpty(t, hotel)
}

func TestCreateHotel(t *testing.T) {
	createHotel(t)
}

func TestGetAllHotels(t *testing.T) {
	createHotel(t)

	result, err := strg.Hotel().GetAll(&repo.GetAllHotelsParams{
		Limit: 3,
		Page:  1,
	})
	require.NoError(t, err)
	require.GreaterOrEqual(t, int(result.Count), 1)
}

func TestUpdateHotel(t *testing.T) {
	c := createHotel(t)

	c.HotelName = faker.NAME
	c.HotelLocation = faker.SENTENCE

	hotel, err := strg.Hotel().Update(c)
	require.NoError(t, err)
	require.NotEmpty(t, hotel)
	require.Equal(t, hotel.HotelName, c.HotelName)
}

func TestDeleteHotel(t *testing.T) {
	c := createHotel(t)

	err := strg.Hotel().Delete(c.ID)
	require.NoError(t, err)
}
