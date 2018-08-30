package seatsvc_test

import (
	"strconv"
	"testing"

	"github.com/nguyenhoaibao/seatsvc"
)

func TestSeatAssign(t *testing.T) {
	var (
		rows        uint = 60
		seatsPerRow uint = 10
	)

	svc := seatsvc.New(rows, seatsPerRow)

	for i := 0; i < 12; i++ {
		if err := svc.Assign(); err != nil {
			t.Fail()
		}
	}
	for i := 1; i < 12; i++ {
		if svc.IsSeatAvailable(uint(i)) {
			t.Errorf("svc.IsSeatAvailable(%d) = true, expected false", i)
		}
	}

	// unassigned seat
	if !svc.IsSeatAvailable(13) {
		t.Error("svc.IsSeatAvailable(13) = false, expected true")
	}
}

func TestSeatName(t *testing.T) {
	var (
		rows        uint = 60
		seatsPerRow uint = 10

		rowName = make([]string, rows)
		colName = make([]string, seatsPerRow)
	)
	for i := 0; i < int(rows); i++ {
		rowName[i] = strconv.Itoa(i + 1)
	}
	for i := 0; i < int(seatsPerRow); i++ {
		colName[i] = string(i + 65)
	}

	svc := seatsvc.New(rows, seatsPerRow)
	svc.SetDimensions(rowName, colName)

	tests := []struct {
		seat uint
		name string
	}{
		{1, "1A"},
		{2, "1B"},
		{10, "1J"},
		{11, "2A"},
		{22, "3B"},
	}
	for _, tt := range tests {
		name := svc.SeatName(tt.seat)
		if name != tt.name {
			t.Errorf("svc.SeatName(%d) = %s, expected %s", tt.seat, name, tt.name)
		}
	}
}
