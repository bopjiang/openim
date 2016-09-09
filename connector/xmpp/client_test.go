package xmpp

import (
	"testing"
)

func TestAdd1(t *testing.T) {
	expect := int(3)
	actual := add(1, 2)
	if actual != expect {
		t.Errorf("expected %d, got %d", expect, actual)
	}
}
