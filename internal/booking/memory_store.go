package booking

type MemoryStore struct {
	bookings map[string]Booking
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: make(map[string]Booking),
	}
}

func (s *MemoryStore) Book(b Booking) (Booking, error) {

	// seat is taken
	if _, ok := s.bookings[b.SeatID]; ok {
		return Booking{}, ErrSeatAlreadyBooked
	}

	// book the seat
	s.bookings[b.SeatID] = b

	return b, nil
}

func (s *MemoryStore) ListBookings(movieID string) []Booking {
	var bookings []Booking
	for _, b := range s.bookings {
		if b.MovieID == movieID {
			bookings = append(bookings, b)
		}
	}
	return bookings
}
