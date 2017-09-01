package backlog

import "testing"

func assertString(t *testing.T, name, actual, expected string) bool {
	if actual != expected {
		t.Errorf("unexpected %v. expected: %v, actual: %v", name, expected, actual)
		return false
	}
	return true
}

func assertInt(t *testing.T, name string, actual, expected int) bool {
	if actual != expected {
		t.Errorf("unexpected %v. expected: %v, actual: %v", name, expected, actual)
		return false
	}
	return true
}
