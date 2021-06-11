package ownmath

import "testing"

const generalError = "value expeted %v, but found value is %v"

func TestMedia(t *testing.T) {
	expeted := 7.38
	value := Media(8, 9.5, 4.65)

	if expeted != value {
		t.Errorf(generalError, expeted, value)
	}
}
