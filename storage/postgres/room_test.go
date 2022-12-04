package postgres_test

import (
	"github.com/MuhammadyusufAdhamov/booking/storage/repo"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRoom(t *testing.T) *repo.Room {
	hotel := createHotel(t)

	room, err := strg.Room().Create(&repo.Room{
		Type:    faker.SENTENCE,
		HotelId: int(hotel.ID),
	})
	require.NoError(t, err)
	require.NotEmpty(t, room)

	return room
}

func TestGetRoom(t *testing.T) {
	c := createRoom(t)

	room, err := strg.Room().Get(c.ID)
	require.NoError(t, err)
	require.NotEmpty(t, room)
}

func TestCreateRoom(t *testing.T) {
	createRoom(t)
}

func TestGetAllRooms(t *testing.T) {
	createRoom(t)

	result, err := strg.Room().GetAll(&repo.GetAllRoomsParams{
		Limit: 3,
		Page:  1,
	})
	require.NoError(t, err)
	require.GreaterOrEqual(t, int(result.Count), 1)
}

func TestUpdateRoom(t *testing.T) {
	c := createRoom(t)

	c.Type = faker.SENTENCE

	room, err := strg.Room().Update(c)
	require.NoError(t, err)
	require.NotEmpty(t, room)
	require.Equal(t, room.Type, c.Type)
}

func TestDeleteRoom(t *testing.T) {
	c := createRoom(t)

	err := strg.Room().Delete(c.ID)
	require.NoError(t, err)
}
