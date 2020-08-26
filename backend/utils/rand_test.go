package utils

import "testing"

func TestRandString(t *testing.T) {
	key := RandString(5)

	t.Errorf("%s", key)

}
func TestRandKey(t *testing.T) {
	key, err := RandKey(32)
	if err == nil {
		t.Logf("%x", key)
	} else {
		t.Error(err)
	}
}
