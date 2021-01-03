package days

import (
	"testing"
)

func TestSeatID(t *testing.T) {
	tests := []struct {
		Row    int
		Column int
		ID     int
	}{
		{1, 5, 13},
		{44, 5, 357},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("SeatID Check", func(t *testing.T) {
			t.Parallel()
			actual := SeatID(tc.Row, tc.Column)
			if actual != tc.ID {
				t.Errorf("got %d, want %d", actual, tc.ID)
			}
		})
	}
}
