package booking

import "sync"

type ConcurrentStore struct {
	bookings map[string]Booking
	sync.RWMutex
}

func NewConcurrentStore() *ConcurrentStore {
	return &ConcurrentStore{
		bookings: make(map[string]Booking),
	}
}

func (s *ConcurrentStore) Book(b Booking) (Booking, error) {

	s.Lock()
	defer s.Unlock()

	// seat is taken
	if _, ok := s.bookings[b.SeatID]; ok {
		return Booking{}, ErrSeatAlreadyBooked
	}

	// book the seat
	s.bookings[b.SeatID] = b

	return b, nil
}

func (s *ConcurrentStore) ListBookings(movieID string) []Booking {

	s.RLock()
	defer s.RUnlock()

	var bookings []Booking
	for _, b := range s.bookings {
		if b.MovieID == movieID {
			bookings = append(bookings, b)
		}
	}
	return bookings
}
