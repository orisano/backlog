package backlog

import (
	"testing"
	"time"
)

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

func assertBool(t *testing.T, name string, actual, expected bool) bool {
	if actual != expected {
		t.Errorf("unexpected %v. expected: %v, actual: %v", name, expected, actual)
		return false
	}
	return true
}

func assertTime(t *testing.T, name string, actual time.Time, year, month, day int) bool {
	ok := true
	ok = assertInt(t, name+"#Year", actual.Year(), year) && ok
	ok = assertInt(t, name+"#Month", int(actual.Month()), month) && ok
	ok = assertInt(t, name+"#Day", actual.Day(), day) && ok
	return ok
}
