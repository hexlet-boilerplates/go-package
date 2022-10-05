package src

import (
	"testing"
)

func TestHello1(t *testing.T) {
	name := "Toto"
	expected := "Hi, Toto. Welcome!"
	actual := Hello(name)

	if actual != expected {
		t.Errorf("Actual %q not equal to expected %q", actual, expected)
	}
}

func TestHello2(t *testing.T) {
	name := "<3"
	expected := "Hi, <3. Welcome!"
	actual := Hello(name)

	if actual != expected {
		t.Errorf("Actual %q not equal to expected %q", actual, expected)
	}
}
