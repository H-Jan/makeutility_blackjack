package math

import "testing"

func TestPlayer(t *testing.T) {
	got := playerScore
	want := (playerScore >= 0)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDealer(t *testing.T) {
	got := dealerScore
	want := (dealerScore >= 0)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
