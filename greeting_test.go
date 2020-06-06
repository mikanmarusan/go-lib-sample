package greeting

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	if "Hello, World" != Greeting() {
		t.Fail()
	}
}
