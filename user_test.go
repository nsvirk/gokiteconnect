package kiteconnect

import (
	"testing"
)

func (ts *TestSuite) TestGetUserProfile(t *testing.T) {
	t.Parallel()
	profile, err := ts.KiteConnect.GetUserProfile()
	if err != nil || profile.Email == "" || profile.UserID == "" {
		t.Errorf("Error while reading user profile. Error: %v", err)
	}
}

func (ts *TestSuite) TestGetFullUserProfile(t *testing.T) {
	t.Parallel()
	fullProfile, err := ts.KiteConnect.GetFullUserProfile()
	if err != nil || fullProfile.Email == "" || fullProfile.UserID == "" {
		t.Errorf("Error while reading full user profile. Error: %v", err)
	}
}

func (ts *TestSuite) TestGetUserMargins(t *testing.T) {
	t.Parallel()
	margins, err := ts.KiteConnect.GetUserMargins()
	if err != nil {
		t.Errorf("Error while reading user margins. Error: %v", err)
	}

	if !margins.Equity.Enabled || !margins.Commodity.Enabled {
		t.Errorf("Incorrect margin values.")
	}
}

func (ts *TestSuite) TestGetUserSegmentMargins(t *testing.T) {
	t.Parallel()
	margins, err := ts.KiteConnect.GetUserSegmentMargins("test")
	if err != nil {
		t.Errorf("Error while reading user margins. Error: %v", err)
	}

	if !margins.Enabled {
		t.Errorf("Incorrect segment margin values.")
	}
}
