package seatsvc

import (
	"errors"
	"fmt"
)

// Service is the seat service.
type Service struct {
	rows             uint // total rows
	cols             uint
	rowName, colName []string // name of the row and col

	maxRowVal uint16 // max number of each row value

	// each element in the seats represents one row in the airplane, it has
	// 16 bit long, and the max number for each element is 1 << seats_per_row.
	// When assigning a seat, it assign from the right most bit first, if
	// that bit is 0, it'll set that to 1, then return. If not, that mean
	// that seat is assigned, so the process will continue to select the next
	// bit, and so on, until it finds a bit that has value is 0.
	// If all the rows are assigned, an error will be returned.
	seats []uint16
}

// New returns the service.
func New(rows, seatsPerRow uint) *Service {
	return &Service{
		rows:      rows,
		cols:      seatsPerRow,
		maxRowVal: 1 << seatsPerRow,
		seats:     make([]uint16, rows),
	}
}

func (s *Service) isRowAvail(i int) bool {
	if s.seats[i] < s.maxRowVal {
		return true
	}
	return false
}

func (s *Service) assignSeat(i int) bool {
	if !s.isRowAvail(i) {
		return false
	}

	var (
		row          = s.seats[i]
		avail uint16 = 1
	)
	for {
		if row&avail == 0 {
			s.seats[i] |= avail
			return true
		}

		avail <<= 1
		if avail >= s.maxRowVal {
			return false
		}
	}
}

// IsSeatAvailable returns true if seat is available.
func (s *Service) IsSeatAvailable(seat uint) bool {
	seat = seat - 1
	var (
		row = s.seats[seat/s.cols]
		col = seat % s.cols
	)
	avail := (row >> col) & 1
	if avail == 1 {
		return false
	}
	return true
}

// Assign finds and assigns available seat.
func (s *Service) Assign() error {
	for i := 0; i < int(s.rows); i++ {
		if s.assignSeat(i) {
			return nil
		}
	}
	return errors.New("no more seats to assign")
}

// SetDimensions sets the row name and col name.
func (s *Service) SetDimensions(rowName, colName []string) error {
	if len(rowName) < int(s.rows) || len(colName) < int(s.cols) {
		return fmt.Errorf("row name must have length %d, col name must have length %d", s.rows, s.cols)
	}
	s.rowName = rowName
	s.colName = colName
	return nil
}

// SeatName returns the seat name based on x and y.
func (s *Service) SeatName(seat uint) string {
	if s.rowName == nil || s.colName == nil {
		return ""
	}

	seat = seat - 1
	var (
		row = seat / s.cols
		col = seat % s.cols
	)
	if int(row) >= len(s.rowName) || int(col) >= len(s.colName) {
		return ""
	}
	return s.rowName[row] + s.colName[col]
}

// Rows return total rows.
func (s *Service) Rows() int {
	return int(s.rows)
}

// Cols return total cols.
func (s *Service) Cols() int {
	return int(s.cols)
}
