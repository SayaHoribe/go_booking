package repository

import "github.com/SayaHoribe/go_booking/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)

	InsertRoomRestriction(r models.RoomRestriction) error
}
