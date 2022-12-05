package postgres_test

import (
	"github.com/MuhammadyusufAdhamov/booking/storage/repo"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
	"testing"
)

func createBooking(t *testing.T) *repo.Booking {
	room := createRoom(t)
	user := createUser(t)
	hotel := createHotel(t)

	booking, err := strg.Booking().Create(&repo.Booking{
		RoomId:   int(room.ID),
		UserId:   int(user.ID),
		HotelId:  int(hotel.ID),
		FromDate: faker.DATE,
	})
	require.NoError(t, err)
	require.NotEmpty(t, booking)

	return booking
}

func TestGetBooking(t *testing.T) {
	c := createBooking(t)

	booking, err := strg.Booking().Get(c.ID)
	require.NoError(t, err)
	require.NotEmpty(t, booking)
}

func TestCreateBooking(t *testing.T) {
	createBooking(t)
}

func TestGetAllBookings(t *testing.T) {
	createBooking(t)

	result, err := strg.Booking().GetAll(&repo.GetAllBookingsParams{
		Limit: 3,
		Page:  1,
	})
	require.NoError(t, err)
	require.GreaterOrEqual(t, int(result.Count), 1)
}

func TestUpdateBooking(t *testing.T) {
	c := createBooking(t)

	c.FromDate = faker.DATE

	booking, err := strg.Booking().Update(c)
	require.NoError(t, err)
	require.NotEmpty(t, booking)
	require.Equal(t, booking.FromDate, c.FromDate)
}

func TestDeleteBooking(t *testing.T) {
	c := createBooking(t)

	err := strg.Booking().Delete(c.ID)
	require.NoError(t, err)
}
