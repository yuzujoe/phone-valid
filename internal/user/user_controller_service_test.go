package user

import (
	"testing"
	"time"
)

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

func Test_checkExpired(t *testing.T) {
	cases := []struct {
		expired time.Time
	}{
		// TODO: Add test cases.
		{expired: time.Now().Add(15 * time.Minute)},
		{expired: time.Now()},
		{expired: time.Now().Add(-15 * time.Minute)},
	}
	for _, tt := range cases {
		if err := checkExpired(tt.expired); err != nil {
			t.Errorf("checkExpired() error = %v", err)
		}
	}
}

func Test_compareCode(t *testing.T) {
	cases := []struct {
		code    string
		reqCode string
	}{
		{code: "187398", reqCode: "187398"},
		{code: "187368", reqCode: "187364"},
	}
	for _, tt := range cases {
		if err := compareCode(tt.code, tt.reqCode); err != nil {
			t.Errorf("compareCode() error = %v", err)
		}
	}
}
