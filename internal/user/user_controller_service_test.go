package user

import "testing"

func Test_phoneValid(t *testing.T) {
	cases := []struct {
		phoneNumber string
	}{
		{phoneNumber: "080-0008-7779"},
		{phoneNumber: "08071985692"},
	}
	for _, tt := range cases {
		if err := phoneValid(tt.phoneNumber); err != nil {
			t.Error(err)
		}
	}
}
